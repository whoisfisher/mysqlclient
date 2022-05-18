package v1alpha1

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type MysqlClusterLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlCluster, err error)
	MysqlCluster(namespace string) MysqlClusterNamespaceLister
	MysqlClusterListerExpansion
}

type mysqlClusterLister struct {
	indexer cache.Indexer
}

func NewMysqlClusterLister(indexer cache.Indexer) MysqlClusterLister {
	return &mysqlClusterLister{indexer: indexer}
}

func (s *mysqlClusterLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlCluster))
	})
	return ret, err
}

func (s *mysqlClusterLister) MysqlCluster(namespace string) MysqlClusterNamespaceLister {
	return mysqlClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type MysqlClusterNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlCluster, err error)
	Get(name string) (*v1alpha1.MysqlCluster, error)
	MysqlClusterNamespaceListerExpansion
}

type mysqlClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s mysqlClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlCluster))
	})
	return ret, err
}

func (s mysqlClusterNamespaceLister) Get(name string) (*v1alpha1.MysqlCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(Resource("MysqlCluster"), name)
	}
	return obj.(*v1alpha1.MysqlCluster), nil
}
