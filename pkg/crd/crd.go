//go:generate go tool gen .
package crd

import (
	"context"

	"github.com/go-json-experiment/json"
	jsonv1 "github.com/go-json-experiment/json/v1"
	"github.com/octohelm/x/ptr"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionstypesv1 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-tools/pkg/crd"
)

type CustomResourceDefinition struct {
	GroupVersion schema.GroupVersion
	KindType     runtime.Object
	ListKindType runtime.Object
	SpecType     any
	Plural       string
	ShortNames   []string
}

func applyCRD(ctx context.Context, apis apiextensionstypesv1.CustomResourceDefinitionInterface, crd *apiextensionsv1.CustomResourceDefinition) error {
	_, err := apis.Get(ctx, crd.Name, metav1.GetOptions{})
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return err
		}
		_, err := apis.Create(ctx, crd, metav1.CreateOptions{})
		return err
	}
	data, err := json.Marshal(crd, jsonv1.OmitEmptyWithLegacySemantics(true))
	if err != nil {
		return err
	}
	_, err = apis.Patch(ctx, crd.Name, types.MergePatchType, data, metav1.PatchOptions{})
	return err
}

func scanJSONSchema(ctx context.Context, v any) *apiextensionsv1.JSONSchemaProps {
	f := &apiextensionsv1.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensionsv1.JSONSchemaProps{
			"spec": {
				XPreserveUnknownFields: ptr.Ptr(true),
			},
			"status": {
				XPreserveUnknownFields: ptr.Ptr(true),
			},
		},
	}

	return crd.FlattenEmbedded(f, nil)
}
