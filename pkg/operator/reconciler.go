package operator

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime"
)

// +gengo:injectable:provider
type ReconcilerRegistry interface {
	RegisterReconciler(reconcilers ...Reconciler) error
}

type Reconciler interface {
	SetupWithManager(mgr controllerruntime.Manager) error
}

type SchemeAdder interface {
	AddToScheme(s *runtime.Scheme) error
}
