package fake

import (
	"github.com/whoisfisher/mysql-client/versioned/typed/mysql/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMysqlV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMysqlV1alpha1) MysqlBackup(namespace string) v1alpha1.MysqlBackupInterface {
	return &FakeMysqlBackup{c, namespace}
}

func (c *FakeMysqlV1alpha1) MysqlCluster(namespace string) v1alpha1.MysqlClusterInterface {
	return &FakeMysqlCluster{c, namespace}
}

func (c *FakeMysqlV1alpha1) MysqlDatabase(namespace string) v1alpha1.MysqlDatabaseInterface {
	return &FakeMysqlDatabase{c, namespace}
}

func (c *FakeMysqlV1alpha1) MysqlUser(namespace string) v1alpha1.MysqlUserInterface {
	return &FakeMysqlUser{c, namespace}
}

func (c *FakeMysqlV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
