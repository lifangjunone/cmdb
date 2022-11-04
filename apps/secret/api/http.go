package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/secret"
	"github.com/lifangjunone/cmdb/service_registry"
)

var (
	h = &handler{}
)

type handler struct {
	service secret.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(h.Name())
	h.service = service_registry.GetGrpcApp(h.Name()).(secret.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return secret.AppName
}

func (h *handler) Version() string {
	return "V1"
}

func (h *handler) Registry(ws *restful.WebService) {
	ws.Route(ws.POST("/").To(h.CreateSecret))
	ws.Route(ws.GET("/").To(h.QuerySecret))
	ws.Route(ws.GET("/{id}").To(h.DescribeSecret))
	ws.Route(ws.DELETE("/{id}").To(h.DeleteSecret))
}

func init() {
	service_registry.RegistryRESTFulApp(h)
}
