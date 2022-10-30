package api

import (
	restFulSpec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/service_registry"

	"github.com/lifangjunone/cmdb/apps/resource"
)

var (
	h = &handler{}
)

type handler struct {
	service resource.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(resource.AppName)
	h.service = service_registry.GetGrpcApp(resource.AppName).(resource.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

// Version 函数定义,  Restful API version
func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	// RESTFul API,   resource = cmdb_resource, action: list, auth: true
	ws.Route(ws.GET("/search").To(h.SearchResource).
		Doc("get all resources").
		Metadata(restFulSpec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Reads(resource.SearchRequest{}).
		Writes(response.NewData(resource.ResourceSet{})).
		Returns(200, "OK", resource.ResourceSet{}))
}

func init() {
	service_registry.RegistryRESTFulApp(h)
}
