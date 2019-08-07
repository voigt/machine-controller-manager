// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	versioned "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/client/listers/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OpenStackMachineClassInformer provides access to a shared informer and lister for
// OpenStackMachineClasses.
type OpenStackMachineClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OpenStackMachineClassLister
}

type openStackMachineClassInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOpenStackMachineClassInformer constructs a new informer for OpenStackMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOpenStackMachineClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOpenStackMachineClassInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOpenStackMachineClassInformer constructs a new informer for OpenStackMachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOpenStackMachineClassInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().OpenStackMachineClasses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().OpenStackMachineClasses(namespace).Watch(options)
			},
		},
		&machinev1alpha1.OpenStackMachineClass{},
		resyncPeriod,
		indexers,
	)
}

func (f *openStackMachineClassInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOpenStackMachineClassInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *openStackMachineClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machinev1alpha1.OpenStackMachineClass{}, f.defaultInformer)
}

func (f *openStackMachineClassInformer) Lister() v1alpha1.OpenStackMachineClassLister {
	return v1alpha1.NewOpenStackMachineClassLister(f.Informer().GetIndexer())
}
