package sample

import (
	"context"

	reconcilersource "github.com/zeiss/knative-extension-sample-source/pkg/reconciler"

	"github.com/zeiss/knative-extension-sample-source/pkg/apis/sources/v1alpha1"

	"github.com/kelseyhightower/envconfig"
	"k8s.io/client-go/tools/cache"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/resolver"

	"github.com/zeiss/knative-extension-sample-source/pkg/reconciler"

	samplesourceinformer "github.com/zeiss/knative-extension-sample-source/pkg/client/injection/informers/sources/v1alpha1/samplesource"
	"github.com/zeiss/knative-extension-sample-source/pkg/client/injection/reconciler/sources/v1alpha1/samplesource"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	deploymentinformer "knative.dev/pkg/client/injection/kube/informers/apps/v1/deployment"
)

// NewController initializes the controller and is called by the generated code
// Registers event handlers to enqueue events
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	deploymentInformer := deploymentinformer.Get(ctx)
	sampleSourceInformer := samplesourceinformer.Get(ctx)

	r := &Reconciler{
		dr: &reconciler.DeploymentReconciler{KubeClientSet: kubeclient.Get(ctx)},
		// Config accessor takes care of tracing/config/logging config propagation to the receive adapter
		configAccessor: reconcilersource.WatchConfigurations(ctx, "sample-source", cmw),
	}
	if err := envconfig.Process("", r); err != nil {
		logging.FromContext(ctx).Panicf("required environment variable is not defined: %v", err)
	}

	impl := samplesource.NewImpl(ctx, r)

	r.sinkResolver = resolver.NewURIResolverFromTracker(ctx, impl.Tracker)

	sampleSourceInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	deploymentInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterController(&v1alpha1.SampleSource{}),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
