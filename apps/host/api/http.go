package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/logger"
	"github.com/lifangjunone/cmdb/apps/host"
)

var (
	h = &handler{}
)

type handler struct {
	service host.ServiceServer
	log     logger.Logger
}

func (h *handler) Name() string {
	return host.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	ws.Route(ws.POST("/").To(h.CreateHost))
	ws.Route(ws.GET("/").To(h.QueryHost))
	ws.Route(ws.GET("/{id}").To(h.DescribeHost))
}
