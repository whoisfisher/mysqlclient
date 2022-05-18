package mysql

import (
	internalinterfaces "github.com/whoisfisher/mysql-client/informers/externalversions/internalinterfaces"
	"github.com/whoisfisher/mysql-client/informers/externalversions/mysql/v1alpha1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	V1alpha1() v1alpha1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

func (g *group) V1alpha1() v1alpha1.Interface {
	return v1alpha1.New(g.factory, g.namespace, g.tweakListOptions)
}
