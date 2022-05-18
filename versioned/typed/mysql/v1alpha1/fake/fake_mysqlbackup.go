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

type FakeMysqlBackup struct {
	Fake *FakeMysqlV1alpha1
	ns   string
}

var MysqlBackupResource = schema.GroupVersionResource{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Resource: "mysqlbackup"}

var MysqlBackupKind = schema.GroupVersionKind{Group: v1alpha1.SchemeGroupVersion.Group, Version: v1alpha1.SchemeGroupVersion.Version, Kind: "mysqlbackup"}

func (c *FakeMysqlBackup) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(MysqlBackupResource, c.ns, name), &v1alpha1.MysqlBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlBackup), err
}

func (c *FakeMysqlBackup) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(MysqlBackupResource, MysqlBackupKind, c.ns, opts), &v1alpha1.MysqlBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlBackupList{ListMeta: obj.(*v1alpha1.MysqlBackupList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *FakeMysqlBackup) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(MysqlBackupResource, c.ns, opts))

}

func (c *FakeMysqlBackup) Create(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts v1.CreateOptions) (result *v1alpha1.MysqlBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(MysqlBackupResource, c.ns, mysqlBackup), &v1alpha1.MysqlBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlBackup), err
}

func (c *FakeMysqlBackup) Update(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts v1.UpdateOptions) (result *v1alpha1.MysqlBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(MysqlBackupResource, c.ns, mysqlBackup), &v1alpha1.MysqlBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlBackup), err
}

func (c *FakeMysqlBackup) UpdateStatus(ctx context.Context, mysqlBackup *v1alpha1.MysqlBackup, opts v1.UpdateOptions) (*v1alpha1.MysqlBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(MysqlBackupResource, "status", c.ns, mysqlBackup), &v1alpha1.MysqlBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlBackup), err
}

func (c *FakeMysqlBackup) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(MysqlBackupResource, c.ns, name), &v1alpha1.MysqlBackup{})

	return err
}

func (c *FakeMysqlBackup) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(MysqlBackupResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlBackup{})
	return err
}

func (c *FakeMysqlBackup) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(MysqlBackupResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlBackup), err
}
