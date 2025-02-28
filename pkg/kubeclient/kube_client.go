package kubeclient

import (
	"context"

	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

// +gengo:injectable:provider
type KubeClient struct {
	// Paths to a kubeconfig. Only required if out-of-cluster.
	Kubeconfig string `flag:",omitempty"`

	c Client `provide:""`
}

func (c *KubeClient) beforeInit(ctx context.Context) error {
	if c.c == nil {
		cc, err := NewClient(c.Kubeconfig)
		if err != nil {
			return err
		}

		if err := clientgoscheme.AddToScheme(cc.Scheme()); err != nil {
			return err
		}

		c.c = cc
	}
	return nil
}
