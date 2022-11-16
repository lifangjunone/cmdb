package impl

import (
	"context"
	"fmt"
	"github.com/lifangjunone/cmdb/apps/resource"
	"github.com/lifangjunone/cmdb/apps/secret"
	"github.com/lifangjunone/cmdb/apps/task"
	"github.com/lifangjunone/cmdb/conf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *impl) CreateTask(ctx context.Context, req *task.CreateTaskRequst) (*task.Task, error) {
	// create task instance
	t, err := task.CreateTask(req)
	if err != nil {
		return nil, err
	}
	// query secret
	sct, err := s.secret.DescribeSecret(ctx, secret.NewDescribeSecretRequest(req.SecretId))
	if err != nil {
		return nil, err
	}
	t.SecretDescription = sct.Data.Description
	return nil, nil
	// decrypt api secret
	if err := sct.Data.DecryptAPISecret(conf.C().App.EncryptKey); err != nil {
		return nil, err
	}
	t.Run()

	var taskCancel context.CancelFunc

	switch req.Type {
	case task.Type_RESOURCE_SYNC:
		switch sct.Data.Vendor {
		case resource.Vendor_TENCENT:
			switch req.ResourceType {
			case resource.Type_HOST:
				taskExecCtx, cancel := context.WithTimeout(
					context.Background(),
					time.Duration(req.Timeout)*time.Second,
				)
				taskCancel = cancel

				go s.syncHost(taskExecCtx, newSyncHostRequest(sct, t))
			case resource.Type_BILL:
			}

		case resource.Vendor_ALIYUN:

		case resource.Vendor_HUAWEI:

		case resource.Vendor_AMAZON:

		case resource.Vendor_VSPHERE:

		default:
			return nil, fmt.Errorf("unknow resource type: %s", sct.Data.Vendor)
		}
	case task.Type_RESOURCE_RELEASE:
	default:
		return nil, fmt.Errorf("unknow task type: %s", req.Type)
	}
	return nil, nil

	if err := s.insert(ctx, t); err != nil {
		if taskCancel != nil {
			taskCancel()
		}
		return nil, err
	}

	return t, nil
}

func (s *impl) QueryTask(ctx context.Context, req *task.QueryTaskRequest) (*task.TaskSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTask not implemented")
}
func (s *impl) DescribeTask(ctx context.Context, req *task.DescribeTaskRequest) (*task.Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTask not implemented")
}
