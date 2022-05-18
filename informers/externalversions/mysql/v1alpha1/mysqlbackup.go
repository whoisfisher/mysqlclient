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

type MysqlBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MysqlBackupLister
}

type mysqlBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func NewMysqlBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMysqlBackupInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewFilteredMysqlBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1alpha1().MysqlBackup(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1alpha1().MysqlBackup(namespace).Watch(context.TODO(), options)
			},
		},
		&mysqlvalpha1.MysqlBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *mysqlBackupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMysqlBackupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *mysqlBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&mysqlvalpha1.MysqlBackup{}, f.defaultInformer)
}

func (f *mysqlBackupInformer) Lister() v1alpha1.MysqlBackupLister {
	return v1alpha1.NewMysqlBackupLister(f.Informer().GetIndexer())
}
