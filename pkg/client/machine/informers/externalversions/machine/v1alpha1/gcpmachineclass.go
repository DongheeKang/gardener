// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/gardener/gardener/pkg/client/machine/clientset/versioned"
	internalinterfaces "github.com/gardener/gardener/pkg/client/machine/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/gardener/pkg/client/machine/listers/machine/v1alpha1"
	machine_v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GCPMachineClassInformer provides access to a shared informer and lister for
// GCPMachineClasses.
type GCPMachineClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.GCPMachineClassLister
}

type gCPMachineClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGCPMachineClassInformer constructs a new informer for GCPMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGCPMachineClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGCPMachineClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGCPMachineClassInformer constructs a new informer for GCPMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGCPMachineClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().GCPMachineClasses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().GCPMachineClasses(namespace).Watch(options)
			},
		},
		&machine_v1alpha1.GCPMachineClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *gCPMachineClassInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGCPMachineClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gCPMachineClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machine_v1alpha1.GCPMachineClass{}, f.defaultInformer)
}

func (f *gCPMachineClassInformer) Lister() v1alpha1.GCPMachineClassLister {
	return v1alpha1.NewGCPMachineClassLister(f.Informer().GetIndexer())
}
