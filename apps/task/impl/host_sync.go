package impl

import (
	"context"
	"fmt"

	"github.com/lifangjunone/cmdb/apps/host"
	"github.com/lifangjunone/cmdb/apps/secret"
	"github.com/lifangjunone/cmdb/apps/task"
	"github.com/lifangjunone/cmdb/provider/tencent/connectivity"
	"github.com/lifangjunone/cmdb/provider/tencent/cvm"
)

func newSyncHostRequest(secret *secret.Secret, task *task.Task) *syncHostReqeust {
	return &syncHostReqeust{
		secret: secret,
		task:   task,
	}
}

type syncHostReqeust struct {
	secret *secret.Secret
	task   *task.Task
}

func (i *impl) syncHost(ctx context.Context, req *syncHostReqeust) {

	// 处理任务状态
	// go routine里面记住 一定要捕获异常, 程序绷掉
	// go recover 只能捕获 当前Gorouine的panice
	defer func() {
		if err := recover(); err != nil {
			// panic 任务失败
			req.task.Failed(fmt.Sprintf("pannic, %v", err))
		} else {
			// 正常结束
			if req.task.Status.Stage != task.Stage_FAILED && req.task.Status.Stage != task.Stage_WARNING {
				req.task.Success()
			}
			if err := i.update(ctx, req.task); err != nil {
				i.log.Errorf("save task error, %s", err)
			}
		}
	}()

	// 只实现主机同步, 初始化腾讯cvm operator
	// NewTencentCloudClient client
	txConn := connectivity.NewTencentCloudClient(
		req.secret.Data.ApiKey,
		req.secret.Data.ApiSecret,
		req.task.Data.Region)

	cvmOp := cvm.NewCVMOperator(txConn.CvmClient())

	// 因为要同步所有资源，需要分页查询
	pagger := cvm.NewPagger(float64(req.secret.Data.RequestRate), cvmOp)
	for pagger.Next() {
		set := host.NewHostSet()
		// 查询分页有错误 反应在Task上面
		if err := pagger.Scan(ctx, set); err != nil {
			req.task.Failed(err.Error())
			return
		}
		// 保持该页数据, 同步时间时, 记录下日志
		for index := range set.Items {
			_, err := i.host.SyncHost(ctx, set.Items[index])
			if err != nil {
				i.log.Errorf("sync host error, %s", err)
				continue
			}
		}
	}
}
