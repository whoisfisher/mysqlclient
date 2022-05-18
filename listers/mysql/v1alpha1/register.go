package v1alpha1

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version}

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	localSchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&v1alpha1.MysqlBackup{},
		&v1alpha1.MysqlBackupList{},
		&v1alpha1.MysqlCluster{},
		&v1alpha1.MysqlClusterList{},
		&v1alpha1.MysqlDatabase{},
		&v1alpha1.MysqlDatabaseList{},
		&v1alpha1.MysqlUser{},
		&v1alpha1.MysqlUserList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
