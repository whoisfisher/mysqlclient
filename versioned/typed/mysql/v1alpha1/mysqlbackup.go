package v1alpha1

import (
	"context"
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	"github.com/whoisfisher/mysql-client/versioned/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"time"
)

type MysqlBackupGetter interface {
	MysqlBackup(namespace string) MysqlBackupInterface
}

type MysqlBackupInterface interface {
	Create(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts metav1.CreateOptions) (*v1alpha1.MysqlBackup, error)
	Update(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts metav1.UpdateOptions) (*v1alpha1.MysqlBackup, error)
	UpdateStatus(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts metav1.UpdateOptions) (*v1alpha1.MysqlBackup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.MysqlBackup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.MysqlBackupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlBackup, err error)
	MysqlBackupExpansion
}

type mysqlBackup struct {
	client rest.Interface
	ns     string
}

func NewMysqlBackup(c *MysqlV1alpha1Client, namespace string) *mysqlBackup {
	return &mysqlBackup{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

func (c *mysqlBackup) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.MysqlBackup, err error) {
	result = &v1alpha1.MysqlBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackup").
		Name(name).
		VersionedParams(&options, schema.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlBackup) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.MysqlBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackup").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlBackup) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlbackup").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *mysqlBackup) Create(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts metav1.CreateOptions) (result *v1alpha1.MysqlBackup, err error) {
	result = &v1alpha1.MysqlBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlbackup").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(mysqlBackup).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlBackup) Update(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts metav1.UpdateOptions) (result *v1alpha1.MysqlBackup, err error) {
	result = &v1alpha1.MysqlBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlbackup").
		Name(mysqlBackup.Name).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(mysqlBackup).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlBackup) UpdateStatus(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts metav1.UpdateOptions) (result *v1alpha1.MysqlBackup, err error) {
	result = &v1alpha1.MysqlBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlbackup").
		Name(mysqlBackup.Name).
		SubResource("status").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(mysqlBackup).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlBackup) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlbackup").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlBackup) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlbackup").
		VersionedParams(&listOpts, schema.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlBackup) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlBackup, err error) {
	result = &v1alpha1.MysqlBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlbackup").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
