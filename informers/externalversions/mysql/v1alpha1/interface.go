package v1alpha1

import (
	internalinterfaces "github.com/whoisfisher/mysql-client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	MysqlBackup() MysqlBackupInformer
	MysqlCluster() MysqlClusterInformer
	MysqlDatabase() MysqlDatabaseInformer
	MysqlUser() MysqlUserInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

func (v *version) MysqlBackup() MysqlBackupInformer {
	return &mysqlBackupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

func (v *version) MysqlCluster() MysqlClusterInformer {
	return &mysqlClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

func (v *version) MysqlDatabase() MysqlDatabaseInformer {
	return &mysqlDatabaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

func (v *version) MysqlUser() MysqlUserInformer {
	return &mysqlUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
