package model

import (
	"context"
	"path/filepath"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/casbin/k8s-authz/pkg/crdadaptor"
	"github.com/casbin/k8s-authz/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type ModelAdaptorLoader struct {
	namespace        string
	isExternalClient bool
	clientset        *versioned.Clientset
}

type ModelAdaptorPair struct {
	Name    string
	Model   model.Model
	Adaptor persist.Adapter
}

func NewModelLoader(namespace string, isExternalClient bool) (*ModelAdaptorLoader, error) {
	var res = &ModelAdaptorLoader{
		namespace:        namespace,
		isExternalClient: isExternalClient,
	}
	var err error
	if isExternalClient {
		err = res.establishExternalClient()
	} else {
		err = res.establishInternalClient()
	}
	return res, err
}

func (m *ModelAdaptorLoader) GetModelAndAdaptors() ([]ModelAdaptorPair, error) {
	list, err := m.clientset.AuthV1().CasbinModels(m.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	res := make([]ModelAdaptorPair, 0)
	for _, crdModel := range list.Items {
		casbinModel, err := model.NewModelFromString(crdModel.Spec.ModelText)
		if err != nil {
			return nil, err
		}
		modelName := crdModel.ObjectMeta.Name
		casbinAdaptor, err := crdadaptor.NewK8sAdaptor(m.namespace, modelName, m.isExternalClient)
		if err != nil {
			return nil, err
		}
		res = append(res, ModelAdaptorPair{
			Model:   casbinModel,
			Adaptor: casbinAdaptor,
			Name:    modelName,
		})
	}
	return res, nil
}

func (m *ModelAdaptorLoader) establishInternalClient() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return err
	}
	m.clientset = clientset
	return nil
}

func (m *ModelAdaptorLoader) establishExternalClient() error {
	home := homedir.HomeDir()
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(home, ".kube", "config"))
	if err != nil {
		return err
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return err
	}
	m.clientset = clientset
	return nil
}
