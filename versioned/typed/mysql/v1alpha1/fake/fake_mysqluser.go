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

type FakeMysqlUser struct {
	Fake *FakeMysqlV1alpha1
	ns   string
}

var MysqlUserResource = schema.GroupVersionResource{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Resource: "MysqlUser"}

var MysqlUserKind = schema.GroupVersionKind{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Kind: "MysqlUser"}

func (c *FakeMysqlUser) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(MysqlUserResource, c.ns, name), &v1alpha1.MysqlUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlUser), err
}

func (c *FakeMysqlUser) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(MysqlUserResource, MysqlUserKind, c.ns, opts), &v1alpha1.MysqlUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlUserList{ListMeta: obj.(*v1alpha1.MysqlUserList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeMysqlUser) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(MysqlUserResource, c.ns, opts))

}

func (c *FakeMysqlUser) Create(ctx context.Context, mysqlUser *v1alpha1.MysqlUser, opts v1.CreateOptions) (result *v1alpha1.MysqlUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(MysqlUserResource, c.ns, mysqlUser), &v1alpha1.MysqlUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlUser), err
}

func (c *FakeMysqlUser) Update(ctx context.Context, mysqlUser *v1alpha1.MysqlUser, opts v1.UpdateOptions) (result *v1alpha1.MysqlUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(MysqlUserResource, c.ns, mysqlUser), &v1alpha1.MysqlUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlUser), err
}

func (c *FakeMysqlUser) UpdateStatus(ctx context.Context, mysqlUser *v1alpha1.MysqlUser, opts v1.UpdateOptions) (*v1alpha1.MysqlUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(MysqlUserResource, "status", c.ns, mysqlUser), &v1alpha1.MysqlUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlUser), err
}

func (c *FakeMysqlUser) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(MysqlUserResource, c.ns, name), &v1alpha1.MysqlUser{})

	return err
}

func (c *FakeMysqlUser) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(MysqlUserResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlUser{})
	return err
}

func (c *FakeMysqlUser) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(MysqlUserResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlUser), err
}
