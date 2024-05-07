package kubeclient

import (
	"os"
	"path"
	"strings"

	contextx "github.com/octohelm/x/context"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewClient(kubeConfigPath string) (client.Client, error) {
	if kubeConfigPath != "" && strings.HasPrefix(kubeConfigPath, "~/") {
		kubeConfigPath = path.Join(os.Getenv("HOME"), kubeConfigPath[2:])
	}

	kubeConfig, err := LoadConfig(kubeConfigPath)
	if err != nil {
		return nil, err
	}

	c, err := client.New(kubeConfig, client.Options{})
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

var Context = contextx.New[client.Client]()

func KubeConfigFromClient(c client.Client) *rest.Config {
	if can, ok := c.(interface{ KubeConfig() *rest.Config }); ok {
		return can.KubeConfig()
	}
	return nil
}
