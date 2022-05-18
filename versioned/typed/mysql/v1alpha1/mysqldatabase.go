package v1alpha1

import (
	"context"
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"github.com/whoisfisher/mysql-client/versioned/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"time"
)

type MysqlDatabaseGetter interface {
	MysqlDatabase(namespace string) MysqlDatabaseInterface
}

type MysqlDatabaseInterface interface {
	Create(ctx context.Context, MysqlDatabase *v1alpha1.MysqlDatabase, opts metav1.CreateOptions) (*v1alpha1.MysqlDatabase, error)
	Update(ctx context.Context, MysqlDatabase *v1alpha1.MysqlDatabase, opts metav1.UpdateOptions) (*v1alpha1.MysqlDatabase, error)
	UpdateStatus(ctx context.Context, MysqlDatabase *v1alpha1.MysqlDatabase, opts metav1.UpdateOptions) (*v1alpha1.MysqlDatabase, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.MysqlDatabase, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.MysqlDatabaseList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlDatabase, err error)
	MysqlDatabaseExpansion
}

type mysqlDatabase struct {
	client rest.Interface
	ns     string
}

func NewMysqlDatabase(c *MysqlV1alpha1Client, namespace string) *mysqlDatabase {
	return &mysqlDatabase{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

func (c *mysqlDatabase) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqldatabase").
		Name(name).
		VersionedParams(&options, schema.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlDatabase) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.MysqlDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqldatabase").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlDatabase) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqldatabase").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *mysqlDatabase) Create(ctx context.Context, MysqlDatabase *v1alpha1.MysqlDatabase, opts metav1.CreateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqldatabase").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlDatabase).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlDatabase) Update(ctx context.Context, MysqlDatabase *v1alpha1.MysqlDatabase, opts metav1.UpdateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqldatabase").
		Name(MysqlDatabase.Name).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlDatabase).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlDatabase) UpdateStatus(ctx context.Context, MysqlDatabase *v1alpha1.MysqlDatabase, opts metav1.UpdateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqldatabase").
		Name(MysqlDatabase.Name).
		SubResource("status").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlDatabase).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlDatabase) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqldatabase").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlDatabase) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqldatabase").
		VersionedParams(&listOpts, schema.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlDatabase) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlDatabase, err error) {
	result = &v1alpha1.MysqlDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqldatabase").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
