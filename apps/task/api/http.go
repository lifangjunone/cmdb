package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/task"
	"github.com/lifangjunone/cmdb/service_registry"
)

var (
	h = &handler{}
)

type handler struct {
	task task.ServiceServer
	log  logger.Logger
}

func (h *handler) Name() string {
	return task.AppName
}

func (h *handler) Config() error {
	h.log = zap.L().Named(h.Name())
	h.task = service_registry.GetGrpcApp(h.Name()).(task.ServiceServer)
	return nil
}

func (h *handler) Version() string {
	return "V1"
}

func (h *handler) Registry(ws *restful.WebService) {
	ws.Route(ws.POST("/").To(h.CreateTask))
	ws.Route(ws.GET("/").To(h.QueryTask))
	ws.Route(ws.GET("/{id}").To(h.DescribeTask))
}

func init() {
	service_registry.RegistryRESTFulApp(h)
}
