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

type MysqlUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MysqlUserLister
}

type mysqlUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func NewMysqlUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMysqlUserInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewFilteredMysqlUserInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1alpha1().MysqlUser(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1alpha1().MysqlUser(namespace).Watch(context.TODO(), options)
			},
		},
		&mysqlvalpha1.MysqlUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *mysqlUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMysqlUserInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mysqlUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mysqlvalpha1.MysqlUser{}, f.defaultInformer)
}

func (f *mysqlUserInformer) Lister() v1alpha1.MysqlUserLister {
	return v1alpha1.NewMysqlUserLister(f.Informer().GetIndexer())
}
