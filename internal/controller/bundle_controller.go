/*
Copyright 2023 VMware Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"fmt"
	apiv1beta2 "github.com/fluxcd/source-controller/api/v1beta2"
	"github.com/garethjevans/bundle-controller/api/v1alpha1"
	"github.com/sirupsen/logrus"
	"github.com/vmware-labs/reconciler-runtime/reconcilers"
	"io"
	"k8s.io/apimachinery/pkg/api/errors"
	"net/http"
	"os"
	"path/filepath"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//+kubebuilder:rbac:groups=source.garethjevans.org,resources=bundles,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=source.garethjevans.org,resources=bundles/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=source.garethjevans.org,resources=bundles/finalizers,verbs=update

func NewBundleReconciler(c reconcilers.Config) *reconcilers.ResourceReconciler[*v1alpha1.Bundle] {
	return &reconcilers.ResourceReconciler[*v1alpha1.Bundle]{
		Name: "Bundle",
		Reconciler: reconcilers.Sequence[*v1alpha1.Bundle]{
			NewMixer(c),
		},
		Config: c,
	}
}

func NewMixer(c reconcilers.Config) reconcilers.SubReconciler[*v1alpha1.Bundle] {
	return &reconcilers.SyncReconciler[*v1alpha1.Bundle]{
		Name: "Mixer",
		Sync: func(ctx context.Context, resource *v1alpha1.Bundle) error {
			// create a temporary directory
			tempDir, err := os.MkdirTemp("", "mixer")
			if err != nil {
				return err
			}

			logrus.Infof("creating temp dir %s", tempDir)

			// for each source that is in the list
			for _, source := range resource.Spec.Include {

				// resolve the input
				key := source.Key(resource.ObjectMeta.Namespace)

				component := GetKind(source.Ref.Kind)

				err := c.Client.Get(ctx, key, component)
				if err != nil {
					if errors.IsNotFound(err) {
						logrus.Warnf("unable to resolve %s", key)
						resource.Status.MarkResourceMissing(key.Name, key.Name, key.Namespace)
					} else {
						logrus.Errorf("error resolving %s", key)
						resource.Status.MarkFailed(err)
					}
					return nil
				}

				// parse the status
				a, ok := component.(Artifacter)
				if !ok {
					logrus.Errorf("component does not have an artifact")
				}

				artifact := a.GetArtifact()
				if artifact != nil {
					logrus.Infof("got artifact %+v", artifact)

					// download the bundle and copy from/to path
					err = DownloadFile(filepath.Join(tempDir, fmt.Sprintf("%s.tar.gz", source.Ref.Name)), artifact.URL)
					if err != nil {
						return err
					}
				}

			}

			return nil
		},
	}
}

func GetKind(kind string) client.Object {
	if kind == "OCIRepository" {
		return &apiv1beta2.OCIRepository{}
	}
	return &apiv1beta2.GitRepository{}
}

type Artifacter interface {
	GetArtifact() *apiv1beta2.Artifact
}

var _ Artifacter = (*apiv1beta2.GitRepository)(nil)
var _ Artifacter = (*apiv1beta2.OCIRepository)(nil)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	logrus.Infof("Downloading %s", url)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func ExtractTarGz(tarGzPath string, dir string) error {
	r, err := os.Open(tarGzPath)
	if err != nil {
		return err
	}

	uncompressedStream, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(header.Name, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
			outFile.Close()

		default:
			return err
		}
	}
	return nil
}
