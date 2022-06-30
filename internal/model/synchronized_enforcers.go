package model

import (
	"fmt"
	"log"
	"sync"
	"time"

	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/k8s-gatekeeper/pkg/casbinhelper"
	admission "k8s.io/api/admission/v1"
)

type EnforcerWrapper struct {
	Enforcer  *casbin.Enforcer
	ModelName string
}

type SynchronizedEnforcerList struct {
	sync.Mutex
	Enforcers []*EnforcerWrapper
	loader    *ModelAdaptorLoader
}

var EnforcerList *SynchronizedEnforcerList

func Init() {
	EnforcerList = NewSynchronizedEnforcerList()
}

func NewSynchronizedEnforcerList() *SynchronizedEnforcerList {
	//todo: switch to dynamic configuration
	loader, err := NewModelLoader("default", IsExternalClient)
	if err != nil {
		panic(err)
	}
	res := &SynchronizedEnforcerList{
		Enforcers: make([]*EnforcerWrapper, 0),
		loader:    loader,
	}
	//load all enabled models and rules
	res.loadEnforcer()
	//start auto sync for loaders
	go func() {
		for {
			time.Sleep(10 * time.Second)
			res.loadEnforcer()
		}
	}()
	return res

}

func (s *SynchronizedEnforcerList) Enforce(admission *admission.AdmissionReview) error {
	s.Lock()
	defer s.Unlock()

	for _, wrapper := range s.Enforcers {
		ok, err := wrapper.Enforcer.Enforce(admission)
		if err != nil {
			return fmt.Errorf("%s rejected the request: %s", wrapper.ModelName, err.Error())
		} else if !ok {
			return fmt.Errorf("%s rejected the request", wrapper.ModelName)
		}
	}
	return nil
}

func (s *SynchronizedEnforcerList) loadEnforcer() {
	s.Lock()
	defer s.Unlock()

	modelAndAdptors, err := s.loader.GetModelAndAdaptors()
	if err != nil {
		log.Printf("error: %s", err.Error())
		return
	}
	s.Enforcers = s.Enforcers[:0]
	for _, tmp := range modelAndAdptors {
		e, err := casbin.NewEnforcer(tmp.Model, tmp.Adaptor)
		if err != nil {
			log.Printf("error: %s", err.Error())
			return
		}
		//todo: setup function lists
		e.AddFunction("access", casbinhelper.Access)
		e.AddFunction("accessWithWildcard", casbinhelper.AccessWithWildCard)
		e.AddFunction("string", casbinhelper.ToString)
		e.AddFunction("parseFloat", casbinhelper.ParseFloat)
		e.AddFunction("contain", casbinhelper.Contain)
		e.AddFunction("split", casbinhelper.Split)
		e.AddFunction("len", casbinhelper.Len)
		e.AddFunction("matchRegex", casbinhelper.MatchRegex)
		e.AddFunction("isNil", casbinhelper.IsNil)
		s.Enforcers = append(s.Enforcers, &EnforcerWrapper{Enforcer: e, ModelName: tmp.Name})
	}
	log.Printf("%d enforcers loaded", len(s.Enforcers))
}
