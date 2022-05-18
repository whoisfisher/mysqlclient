package externalversions

import (
	"fmt"
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	case v1alpha1.SchemeGroupVersion.WithResource("mysqlbackup"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mysql().V1alpha1().MysqlBackup().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("mysqlcluster"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mysql().V1alpha1().MysqlCluster().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("mysqldatabase"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mysql().V1alpha1().MysqlDatabase().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("mysqluser"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Mysql().V1alpha1().MysqlUser().Informer()}, nil
	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
