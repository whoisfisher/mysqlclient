package v1alpha1

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type MysqlBackupLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlBackup, err error)
	MysqlBackup(namespace string) MysqlBackupNamespaceLister
	MysqlBackupListerExpansion
}

type mysqlBackupLister struct {
	indexer cache.Indexer
}

func NewMysqlBackupLister(indexer cache.Indexer) MysqlBackupLister {
	return &mysqlBackupLister{indexer: indexer}
}

func (s *mysqlBackupLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlBackup))
	})
	return ret, err
}

func (s *mysqlBackupLister) MysqlBackup(namespace string) MysqlBackupNamespaceLister {
	return mysqlBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type MysqlBackupNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.MysqlBackup, err error)
	Get(name string) (*v1alpha1.MysqlBackup, error)
	MysqlBackupNamespaceListerExpansion
}

type mysqlBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s mysqlBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlBackup))
	})
	return ret, err
}

func (s mysqlBackupNamespaceLister) Get(name string) (*v1alpha1.MysqlBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(Resource("mysqlbackup"), name)
	}
	return obj.(*v1alpha1.MysqlBackup), nil
}
