package operator

import (
	contextx "github.com/octohelm/x/context"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime"
)

type Reconciler interface {
	SetupWithManager(mgr controllerruntime.Manager) error
}

type ReconcilerRegistry interface {
	RegisterReconciler(reconcilers ...Reconciler) error
}

var ReconcilerRegistryContext = contextx.New[ReconcilerRegistry]()

type SchemeAdder interface {
	AddToScheme(s *runtime.Scheme) error
}
