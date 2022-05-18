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

type FakeMysqlDatabase struct {
	Fake *FakeMysqlV1alpha1
	ns   string
}

var MysqlDatabaseResource = schema.GroupVersionResource{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Resource: "MysqlDatabase"}

var MysqlDatabaseKind = schema.GroupVersionKind{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Kind: "MysqlDatabase"}

func (c *FakeMysqlDatabase) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(MysqlDatabaseResource, c.ns, name), &v1alpha1.MysqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

func (c *FakeMysqlDatabase) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(MysqlDatabaseResource, MysqlDatabaseKind, c.ns, opts), &v1alpha1.MysqlDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlDatabaseList{ListMeta: obj.(*v1alpha1.MysqlDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeMysqlDatabase) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(MysqlDatabaseResource, c.ns, opts))

}

func (c *FakeMysqlDatabase) Create(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.CreateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(MysqlDatabaseResource, c.ns, mysqlDatabase), &v1alpha1.MysqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

func (c *FakeMysqlDatabase) Update(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.UpdateOptions) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(MysqlDatabaseResource, c.ns, mysqlDatabase), &v1alpha1.MysqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

func (c *FakeMysqlDatabase) UpdateStatus(ctx context.Context, mysqlDatabase *v1alpha1.MysqlDatabase, opts v1.UpdateOptions) (*v1alpha1.MysqlDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(MysqlDatabaseResource, "status", c.ns, mysqlDatabase), &v1alpha1.MysqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}

func (c *FakeMysqlDatabase) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(MysqlDatabaseResource, c.ns, name), &v1alpha1.MysqlDatabase{})

	return err
}

func (c *FakeMysqlDatabase) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(MysqlDatabaseResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlDatabase{})
	return err
}

func (c *FakeMysqlDatabase) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(MysqlDatabaseResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlDatabase), err
}
