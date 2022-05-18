package v1alpha1

import (
	"context"
	"github.com/whoisfisher/mysql-client/listers/mysql/v1alpha1"
	time "time"

	mysqlvalpha1 "github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	internalinterfaces "github.com/whoisfisher/mysql-client/informers/externalversions/internalinterfaces"
	versioned "github.com/whoisfisher/mysql-client/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

type MysqlClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MysqlClusterLister
}

type mysqlClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func NewMysqlClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMysqlClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewFilteredMysqlClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1alpha1().MysqlCluster(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1alpha1().MysqlCluster(namespace).Watch(context.TODO(), options)
			},
		},
		&mysqlvalpha1.MysqlCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *mysqlClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMysqlClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mysqlClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mysqlvalpha1.MysqlCluster{}, f.defaultInformer)
}

func (f *mysqlClusterInformer) Lister() v1alpha1.MysqlClusterLister {
	return v1alpha1.NewMysqlClusterLister(f.Informer().GetIndexer())
}
