package client

import (
	"github.com/casbin/k8s-gatekeeper/pkg/generated/clientset/versioned"
)

type K8sGateKeeperClient struct {
	namespace string
	modelName string
	Clientset *versioned.Clientset
}

func NewK8sGateKeeperClient(externalClient bool) (*K8sGateKeeperClient, error) {
	res := K8sGateKeeperClient{
		namespace: "",
		modelName: "",
	}
	var err error
	if externalClient {
		err = res.establishExternalClient()
	} else {
		err = res.establishInternalClient()
	}
	if err != nil {
		return nil, err
	}
	return &res, nil

}

func (k *K8sGateKeeperClient) Namespace(namespace string) *K8sGateKeeperClient {
	k.namespace = namespace
	return k
}

func (k *K8sGateKeeperClient) ModelName(modelName string) *K8sGateKeeperClient {
	k.modelName = modelName
	return k
}
