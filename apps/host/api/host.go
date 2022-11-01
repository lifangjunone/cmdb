package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/lifangjunone/cmdb/apps/host"
)

func (h *handler) CreateHost(r *restful.Request, w *restful.Response) {
	ins := host.NewDefaultHost()
	if err := request.GetDataFromRequest(r.Request, ins); err != nil {
		response.Failed(w, err)
		return
	}
	ins, err := h.service.SyncHost(r.Request.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) QueryHost(r *restful.Request, w *restful.Response) {
	query := host.NewQueryHostRequestFromHttp(r.Request)
	set, err := h.service.QueryHost(r.Request.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeHost(r *restful.Request, w *restful.Response) {
	req := host.NewDescribeHostRequestWithID(r.PathParameter("id"))
	set, err := h.service.DescribeHost(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
