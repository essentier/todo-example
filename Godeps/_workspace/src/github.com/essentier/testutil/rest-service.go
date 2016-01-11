package testutil

import (
	"testing"

	"github.com/essentier/gopencils"
	"github.com/essentier/spickspan"
	"github.com/essentier/spickspan/config"
	"github.com/essentier/spickspan/model"
)

var provider model.Provider

func init() {
	config, err := config.GetConfig()
	if err != nil {
		panic("Failed to find and parse spickspan.json. The error is " + err.Error())
	}

	provider, err = spickspan.GetNomockProvider(config)
	if err != nil {
		panic("Failed to get nomock provider. The error is " + err.Error())
	}

	err = spickspan.BuildAllInConfig(config)
	if err != nil {
		panic("Failed to build projects. The error is " + err.Error())
	}
}

func CreateRestService(serviceName string, t *testing.T) *RestService {
	service, err := provider.GetService(serviceName)
	if err != nil {
		t.Fatalf("Failed to create service %v. Error is: %v.", serviceName, err)
	}

	errHandler := &failTestRestErrHanlder{t: t}
	api := gopencils.Api(service.GetUrl())
	rw := &ResWrapper{Resource: api, ErrHandler: errHandler}
	return &RestService{provider: provider, Service: service, Api: rw}
}

type RestService struct {
	Api      *ResWrapper
	provider model.Provider
	Service  model.Service
}

func (s *RestService) Release() {
	s.provider.Release(s.Service)
}

func (s *RestService) Resource(resourceName string) *ResWrapper {
	return s.Api.NewChildResource(resourceName)
}
