package task

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
	"time"
)

const (
	AppName = "task"
)

var (
	validate = validator.New()
)

func (x *CreateTaskRequst) Validate() error {
	return validate.Struct(x)
}

func CreateTask(req *CreateTaskRequst) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ins := NewDefaultTask()
	ins.Data = req
	ins.Id = xid.New().String()

	return ins, nil
}

func NewDefaultTask() *Task {
	return &Task{
		Data:   &CreateTaskRequst{},
		Status: &Status{},
	}
}

func (s *Task) Run() {
	s.Status.StartAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_RUNNING
}

func (s *Task) Failed(message string) {
	s.Status.EndAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_FAILED
	s.Status.Message = message
}

func (s *Task) Success() {
	s.Status.EndAt = time.Now().UnixMilli()
	s.Status.Stage = Stage_SUCCESS
}

func NewCreateTaskRequest() *CreateTaskRequst {
	return &CreateTaskRequst{
		Params:  map[string]string{},
		Timeout: 30 * 60,
	}
}
