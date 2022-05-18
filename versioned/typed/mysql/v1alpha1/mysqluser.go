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

type MysqlUserGetter interface {
	MysqlUser(namespace string) MysqlUserInterface
}

type MysqlUserInterface interface {
	Create(ctx context.Context, MysqlUser *v1alpha1.MysqlUser, opts metav1.CreateOptions) (*v1alpha1.MysqlUser, error)
	Update(ctx context.Context, MysqlUser *v1alpha1.MysqlUser, opts metav1.UpdateOptions) (*v1alpha1.MysqlUser, error)
	UpdateStatus(ctx context.Context, MysqlUser *v1alpha1.MysqlUser, opts metav1.UpdateOptions) (*v1alpha1.MysqlUser, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.MysqlUser, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.MysqlUserList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlUser, err error)
	MysqlUserExpansion
}

type mysqlUser struct {
	client rest.Interface
	ns     string
}

func NewMysqlUser(c *MysqlV1alpha1Client, namespace string) *mysqlUser {
	return &mysqlUser{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

func (c *mysqlUser) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.MysqlUser, err error) {
	result = &v1alpha1.MysqlUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqluser").
		Name(name).
		VersionedParams(&options, schema.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlUser) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.MysqlUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqluser").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlUser) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqluser").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *mysqlUser) Create(ctx context.Context, MysqlUser *v1alpha1.MysqlUser, opts metav1.CreateOptions) (result *v1alpha1.MysqlUser, err error) {
	result = &v1alpha1.MysqlUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqluser").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlUser).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlUser) Update(ctx context.Context, MysqlUser *v1alpha1.MysqlUser, opts metav1.UpdateOptions) (result *v1alpha1.MysqlUser, err error) {
	result = &v1alpha1.MysqlUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqluser").
		Name(MysqlUser.Name).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlUser).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlUser) UpdateStatus(ctx context.Context, MysqlUser *v1alpha1.MysqlUser, opts metav1.UpdateOptions) (result *v1alpha1.MysqlUser, err error) {
	result = &v1alpha1.MysqlUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqluser").
		Name(MysqlUser.Name).
		SubResource("status").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlUser).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlUser) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqluser").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlUser) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqluser").
		VersionedParams(&listOpts, schema.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlUser) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlUser, err error) {
	result = &v1alpha1.MysqlUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqluser").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
