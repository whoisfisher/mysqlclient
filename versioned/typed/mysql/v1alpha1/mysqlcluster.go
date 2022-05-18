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

type MysqlClusterGetter interface {
	MysqlCluster(namespace string) MysqlClusterInterface
}

type MysqlClusterInterface interface {
	Create(ctx context.Context, MysqlCluster *v1alpha1.MysqlCluster, opts metav1.CreateOptions) (*v1alpha1.MysqlCluster, error)
	Update(ctx context.Context, MysqlCluster *v1alpha1.MysqlCluster, opts metav1.UpdateOptions) (*v1alpha1.MysqlCluster, error)
	UpdateStatus(ctx context.Context, MysqlCluster *v1alpha1.MysqlCluster, opts metav1.UpdateOptions) (*v1alpha1.MysqlCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha1.MysqlCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha1.MysqlClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlCluster, err error)
	MysqlClusterExpansion
}

type mysqlCluster struct {
	client rest.Interface
	ns     string
}

func NewMysqlCluster(c *MysqlV1alpha1Client, namespace string) *mysqlCluster {
	return &mysqlCluster{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

func (c *mysqlCluster) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha1.MysqlCluster, err error) {
	result = &v1alpha1.MysqlCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlcluster").
		Name(name).
		VersionedParams(&options, schema.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlCluster) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha1.MysqlClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MysqlClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlcluster").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlCluster) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlcluster").
		VersionedParams(&opts, schema.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

func (c *mysqlCluster) Create(ctx context.Context, MysqlCluster *v1alpha1.MysqlCluster, opts metav1.CreateOptions) (result *v1alpha1.MysqlCluster, err error) {
	result = &v1alpha1.MysqlCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlcluster").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlCluster).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlCluster) Update(ctx context.Context, MysqlCluster *v1alpha1.MysqlCluster, opts metav1.UpdateOptions) (result *v1alpha1.MysqlCluster, err error) {
	result = &v1alpha1.MysqlCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlcluster").
		Name(MysqlCluster.Name).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlCluster).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlCluster) UpdateStatus(ctx context.Context, MysqlCluster *v1alpha1.MysqlCluster, opts metav1.UpdateOptions) (result *v1alpha1.MysqlCluster, err error) {
	result = &v1alpha1.MysqlCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlcluster").
		Name(MysqlCluster.Name).
		SubResource("status").
		VersionedParams(&opts, schema.ParameterCodec).
		Body(MysqlCluster).
		Do(ctx).
		Into(result)
	return
}

func (c *mysqlCluster) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlcluster").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlCluster) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlcluster").
		VersionedParams(&listOpts, schema.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

func (c *mysqlCluster) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlCluster, err error) {
	result = &v1alpha1.MysqlCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlcluster").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, schema.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
