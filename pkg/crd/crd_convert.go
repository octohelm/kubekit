package crd

import (
	"context"
	"reflect"
	"strings"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func AsKubeCRD(d *CustomResourceDefinition) *apiextensionsv1.CustomResourceDefinition {
	c := &apiextensionsv1.CustomResourceDefinition{}

	kindType := reflect.Indirect(reflect.ValueOf(d.KindType)).Type()

	crdNames := apiextensionsv1.CustomResourceDefinitionNames{
		Kind:       kindType.Name(),
		ListKind:   reflect.Indirect(reflect.ValueOf(d.ListKindType)).Type().Name(),
		ShortNames: d.ShortNames,
	}

	crdNames.Singular = strings.ToLower(crdNames.Kind)

	if d.Plural != "" {
		crdNames.Plural = d.Plural
	} else {
		crdNames.Plural = crdNames.Singular + "s"
	}

	c.Name = crdNames.Plural + "." + d.GroupVersion.Group
	c.Spec.Group = d.GroupVersion.Group
	c.Spec.Scope = apiextensionsv1.NamespaceScoped

	openapiSchema := scanJSONSchema(context.Background(), d.SpecType)

	c.Spec.Names = crdNames
	c.Spec.Versions = []apiextensionsv1.CustomResourceDefinitionVersion{
		{
			Name:    d.GroupVersion.Version,
			Served:  true,
			Storage: true,
			Schema: &apiextensionsv1.CustomResourceValidation{
				OpenAPIV3Schema: openapiSchema,
			},
			Subresources: &apiextensionsv1.CustomResourceSubresources{
				Status: &apiextensionsv1.CustomResourceSubresourceStatus{},
			},
		},
	}

	return c
}
