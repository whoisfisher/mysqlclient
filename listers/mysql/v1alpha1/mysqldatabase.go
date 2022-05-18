package v1alpha1

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type MysqlDatabaseLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlDatabase, err error)
	MysqlDatabase(namespace string) MysqlDatabaseNamespaceLister
	MysqlDatabaseListerExpansion
}

type mysqlDatabaseLister struct {
	indexer cache.Indexer
}

func NewMysqlDatabaseLister(indexer cache.Indexer) MysqlDatabaseLister {
	return &mysqlDatabaseLister{indexer: indexer}
}

func (s *mysqlDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlDatabase))
	})
	return ret, err
}

func (s *mysqlDatabaseLister) MysqlDatabase(namespace string) MysqlDatabaseNamespaceLister {
	return mysqlDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type MysqlDatabaseNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlDatabase, err error)
	Get(name string) (*v1alpha1.MysqlDatabase, error)
	MysqlDatabaseNamespaceListerExpansion
}

type mysqlDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s mysqlDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlDatabase))
	})
	return ret, err
}

func (s mysqlDatabaseNamespaceLister) Get(name string) (*v1alpha1.MysqlDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(Resource("MysqlDatabase"), name)
	}
	return obj.(*v1alpha1.MysqlDatabase), nil
}
