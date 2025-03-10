package crd

import (
	"context"
	"fmt"

	"github.com/octohelm/kubekit/pkg/kubeclient"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Apply(ctx context.Context, c client.Client, crds ...*apiextensionsv1.CustomResourceDefinition) error {
	kubeconfig := kubeclient.KubeConfigFromClient(c)
	if kubeconfig == nil {
		return fmt.Errorf("missing kubeconfig of %T", c)
	}
	cs, err := apiextensionsclientset.NewForConfig(kubeconfig)
	if err != nil {
		return err
	}

	apis := cs.ApiextensionsV1().CustomResourceDefinitions()

	for i := range crds {
		if err := applyCRD(ctx, apis, crds[i]); err != nil {
			return err
		}
	}

	return nil
}
