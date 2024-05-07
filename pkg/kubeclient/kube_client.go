package kubeclient

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeClient struct {
	// Paths to a kubeconfig. Only required if out-of-cluster.
	Kubeconfig string `flag:",omitempty"`

	c client.Client
}

func (c *KubeClient) Init(ctx context.Context) error {
	if c.c == nil {
		cc, err := NewClient(c.Kubeconfig)
		if err != nil {
			return err
		}
		c.c = cc
	}
	return nil
}

func (c *KubeClient) InjectContext(ctx context.Context) context.Context {
	return Context.Inject(ctx, c.c)
}
