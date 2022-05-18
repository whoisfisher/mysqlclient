package fake

import (
	"context"
	"github.com/bitpoke/mysql-operator/pkg/apis/mysql/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeMysqlCluster struct {
	Fake *FakeMysqlV1alpha1
	ns   string
}

var MysqlClusterResource = schema.GroupVersionResource{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Resource: "MysqlCluster"}

var MysqlClusterKind = schema.GroupVersionKind{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Kind: "MysqlCluster"}

func (c *FakeMysqlCluster) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(MysqlClusterResource, c.ns, name), &v1alpha1.MysqlCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlCluster), err
}

func (c *FakeMysqlCluster) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(MysqlClusterResource, MysqlClusterKind, c.ns, opts), &v1alpha1.MysqlClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlClusterList{ListMeta: obj.(*v1alpha1.MysqlClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeMysqlCluster) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(MysqlClusterResource, c.ns, opts))

}

func (c *FakeMysqlCluster) Create(ctx context.Context, mysqlCluster *v1alpha1.MysqlCluster, opts v1.CreateOptions) (result *v1alpha1.MysqlCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(MysqlClusterResource, c.ns, mysqlCluster), &v1alpha1.MysqlCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlCluster), err
}

func (c *FakeMysqlCluster) Update(ctx context.Context, mysqlCluster *v1alpha1.MysqlCluster, opts v1.UpdateOptions) (result *v1alpha1.MysqlCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(MysqlClusterResource, c.ns, mysqlCluster), &v1alpha1.MysqlCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlCluster), err
}

func (c *FakeMysqlCluster) UpdateStatus(ctx context.Context, mysqlCluster *v1alpha1.MysqlCluster, opts v1.UpdateOptions) (*v1alpha1.MysqlCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(MysqlClusterResource, "status", c.ns, mysqlCluster), &v1alpha1.MysqlCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlCluster), err
}

func (c *FakeMysqlCluster) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(MysqlClusterResource, c.ns, name), &v1alpha1.MysqlCluster{})

	return err
}

func (c *FakeMysqlCluster) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(MysqlClusterResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlCluster{})
	return err
}

func (c *FakeMysqlCluster) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(MysqlClusterResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlCluster), err
}
