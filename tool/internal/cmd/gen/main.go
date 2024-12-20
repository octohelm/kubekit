package main

import (
	"context"

	"github.com/octohelm/gengo/pkg/gengo"

	_ "github.com/octohelm/courier/devpkg/clientgen"
	_ "github.com/octohelm/courier/devpkg/injectablegen"
	_ "github.com/octohelm/courier/devpkg/operatorgen"
	_ "github.com/octohelm/gengo/devpkg/deepcopygen"
	_ "github.com/octohelm/gengo/devpkg/runtimedocgen"
)

func main() {
	c, err := gengo.NewContext(&gengo.GeneratorArgs{
		Entrypoint: []string{
			"github.com/octohelm/kubekit/pkg/crd",
			"github.com/octohelm/kubekit/pkg/metadata",
			"github.com/octohelm/kubekit/pkg/operator",
		},
		OutputFileBaseName: "zz_generated",
		Globals: map[string][]string{
			"gengo:runtimedoc": {},
		},
	})
	if err != nil {
		panic(err)
	}

	if err := c.Execute(context.Background(), gengo.GetRegisteredGenerators()...); err != nil {
		panic(err)
	}
}
