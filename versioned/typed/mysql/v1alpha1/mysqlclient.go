package v1alpha1

import (
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"github.com/whoisfisher/mysql-client/versioned/schema"
	"k8s.io/client-go/rest"
)

type MysqlV1alpha1Interface interface {
	RESTClient() rest.Interface
	MysqlBackupGetter
	MysqlClusterGetter
	MysqlDatabaseGetter
	MysqlUserGetter
}

type MysqlV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MysqlV1alpha1Client) MysqlBackup(namespace string) MysqlBackupInterface {
	return NewMysqlBackup(c, namespace)
}

func (c *MysqlV1alpha1Client) MysqlCluster(namespace string) MysqlClusterInterface {
	return NewMysqlCluster(c, namespace)
}

func (c *MysqlV1alpha1Client) MysqlDatabase(namespace string) MysqlDatabaseInterface {
	return NewMysqlDatabase(c, namespace)
}

func (c *MysqlV1alpha1Client) MysqlUser(namespace string) MysqlUserInterface {
	return NewMysqlUser(c, namespace)
}

func NewForConfig(c *rest.Config) (*MysqlV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MysqlV1alpha1Client{client}, nil
}

func NewForConfigOrDie(c *rest.Config) *MysqlV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

func New(c rest.Interface) *MysqlV1alpha1Client {
	return &MysqlV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = schema.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MysqlV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
