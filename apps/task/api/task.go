package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/lifangjunone/cmdb/apps/task"
)

func (h *handler) CreateTask(r *restful.Request, w *restful.Response) {
	req := task.NewCreateTaskRequest()
	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}
	set, err := h.task.CreateTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) QueryTask(r *restful.Request, w *restful.Response) {
	response.Success(w, nil)
}

func (h *handler) DescribeTask(r *restful.Request, w *restful.Response) {
	response.Success(w, nil)
}
