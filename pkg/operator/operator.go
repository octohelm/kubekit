package operator

import (
	"context"
	"fmt"

	"github.com/go-courier/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"

	"github.com/octohelm/kubekit/pkg/kubeclient"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	controllerruntime "sigs.k8s.io/controller-runtime"
)

// +gengo:injectable:provider ReconcilerRegistry
type Operator struct {
	// Watch namespace
	WatchNamespace string `flag:",omitempty"`
	// The address the metric endpoint binds to
	MetricsAddr string `flag:",omitempty"`
	// enable leader election when LeaderElectionID not empty
	LeaderElectionID string `flag:",omitempty"`

	reconcilers []Reconciler `flag:"-"`
	cancel      func()

	client kubeclient.Client `inject:""`
}

func (s *Operator) RegisterReconciler(reconcilers ...Reconciler) error {
	s.reconcilers = append(s.reconcilers, reconcilers...)
	return nil
}

func (s *Operator) Serve(ctx context.Context) error {
	scheme := runtime.NewScheme()

	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		return err
	}

	for _, r := range s.reconcilers {
		if a, ok := r.(SchemeAdder); ok {
			if err := a.AddToScheme(scheme); err != nil {
				return fmt.Errorf("unable to add to scheme %T: %w", a, err)
			}
		}
	}

	ctrlOpt := controllerruntime.Options{
		Logger:           wrapAsGoLogger(logr.FromContext(ctx)),
		Scheme:           scheme,
		LeaderElectionID: s.LeaderElectionID,
		LeaderElection:   s.LeaderElectionID != "",
	}

	ctrlOpt.Metrics.BindAddress = s.MetricsAddr

	if s.WatchNamespace != "" {
		ctrlOpt.Cache.DefaultNamespaces = map[string]cache.Config{
			s.WatchNamespace: {},
		}
	}

	rawClient := kubeclient.KubeConfigFromClient(s.client)

	mgr, err := controllerruntime.NewManager(rawClient, ctrlOpt)
	if err != nil {
		return err
	}

	for _, r := range s.reconcilers {
		if err := r.SetupWithManager(mgr); err != nil {
			return fmt.Errorf("unable to create controller: %T: %w", r, err)
		}
	}

	cc, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	return mgr.Start(cc)
}

func (o *Operator) Shutdown(ctx context.Context) error {
	if o.cancel != nil {
		o.cancel()
	}
	return nil
}
