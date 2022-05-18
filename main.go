package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/whoisfisher/mysql-client/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type Host string

type Config struct {
	Hosts []Host
	Token string
}

func NewMysqlClient(c *Config) (*versioned.Clientset, error) {
	var aliveHost Host
	aliveHost = "192.168.111.100:6443"
	kubeConf := &rest.Config{
		Host:        string(aliveHost),
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}
	client, err := versioned.NewForConfig(kubeConf)
	if err != nil {
		return client, errors.Wrap(err, fmt.Sprintf("new monitoring client with config failed: %v", err))
	}
	return client, nil
}

func getMysqlClient(clusterName string) (*versioned.Clientset, error) {
	var client *versioned.Clientset
	client, err := NewMysqlClient(&Config{
		Token: "eyJhbGciOiJSUzI1NiIsImtpZCI6IlJZeG9KWmtXYnc0X2NUZ2p1U2VmdU9YRS1ZcjVJZHk2QS1qd3BxMWllME0ifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJrby1hZG1pbi10b2tlbi1wZ2toNSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJrby1hZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImZlOGYxY2I3LTJjY2ItNGE2OS1hMzZmLWQ5NzNlMTczYzBiMyIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTprby1hZG1pbiJ9.WemsF5j5JXwRavq78H3qVrble6jSERbvxFxNrgWnJ9iVQSfoQHkpEoisTjOGl_vbg7DecWt5IPbQnW4XVsVfMzQ0wgbP65VHjI2pzx9bfIizCY53shVXUFqVf5H4PCMG_e-ofiq2NRmCkSRczKERIbHV0hqeApUd273EJyjXow_enm6XAueGMAQOc7wCln38NCdrZnRe09J5r96jy6wSAN2bwHoT3z3PXwpbSTLxlakN_kQBKh99HkLau3BhdohCUpKk0HKCRfkMvXZL1tL0s1eJaq6MFgdx9H9NEaE341jc_y0q_cWYJXj7ejbGEHfkn_rDUo47NQbCqK2Huhh6cw",
	})
	if err != nil {
		return client, err
	}
	return client, nil
}

func main() {
	client, err := getMysqlClient("mm")
	if err != nil {
		errMsg := fmt.Sprintf("====%s", err.Error())
		fmt.Println(errMsg)
	}
	snps, err := client.MysqlV1alpha1().MysqlCluster("parent-a").Get(context.TODO(), "child-b", metav1.GetOptions{})
	if err != nil {
		errMsg := fmt.Sprintf("====%s", err.Error())
		fmt.Println(errMsg)
	}
	res := fmt.Sprintf("***%v", snps)
	fmt.Println(res)
}
