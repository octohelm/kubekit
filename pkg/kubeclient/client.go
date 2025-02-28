//go:generate go tool gen .
package kubeclient

import (
	"os"
	"path"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +gengo:injectable:provider
type Client = client.Client

func NewClient(kubeConfigPath string) (Client, error) {
	if kubeConfigPath != "" && strings.HasPrefix(kubeConfigPath, "~/") {
		kubeConfigPath = path.Join(os.Getenv("HOME"), kubeConfigPath[2:])
	}

	kubeConfig, err := LoadConfig(kubeConfigPath)
	if err != nil {
		return nil, err
	}

	c, err := client.New(kubeConfig, client.Options{Scheme: runtime.NewScheme()})
	if err != nil {
		return nil, err
	}

	return &clientWithConfig{Client: c, config: kubeConfig}, nil
}

type clientWithConfig struct {
	client.Client
	config *rest.Config
}

func (c *clientWithConfig) KubeConfig() *rest.Config {
	return c.config
}

func KubeConfigFromClient(c client.Client) *rest.Config {
	if can, ok := c.(interface{ KubeConfig() *rest.Config }); ok {
		return can.KubeConfig()
	}
	return nil
}
