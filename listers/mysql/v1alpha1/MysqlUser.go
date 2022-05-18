package v1alpha1

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type MysqlUserLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlUser, err error)
	MysqlUser(namespace string) MysqlUserNamespaceLister
	MysqlUserListerExpansion
}

type mysqlUserLister struct {
	indexer cache.Indexer
}

func NewMysqlUserLister(indexer cache.Indexer) MysqlUserLister {
	return &mysqlUserLister{indexer: indexer}
}

func (s *mysqlUserLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlUser))
	})
	return ret, err
}

func (s *mysqlUserLister) MysqlUser(namespace string) MysqlUserNamespaceLister {
	return mysqlUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type MysqlUserNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlUser, err error)
	Get(name string) (*v1alpha1.MysqlUser, error)
	MysqlUserNamespaceListerExpansion
}

type mysqlUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s mysqlUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlUser))
	})
	return ret, err
}

func (s mysqlUserNamespaceLister) Get(name string) (*v1alpha1.MysqlUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(Resource("MysqlUser"), name)
	}
	return obj.(*v1alpha1.MysqlUser), nil
}
