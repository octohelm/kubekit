package operator

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime"
)

// +gengo:injectable:provider
type ReconcilerRegistry interface {
	RegisterReconciler(reconcilers ...Reconciler) error
}

type SchemeAdder interface {
	AddToScheme(s *runtime.Scheme) error
}

type Reconciler interface {
	SetupWithManager(mgr controllerruntime.Manager) error
}

type ReconcilerFunc func(mgr controllerruntime.Manager) error

func (fn ReconcilerFunc) SetupWithManager(mgr controllerruntime.Manager) error {
	return fn(mgr)
}
