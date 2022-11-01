package cvm

import (
	"github.com/lifangjunone/cmdb/apps/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type pager struct {
	rep     *cvm.DescribeInstancesRequest
	op      *CVMOperator
	hasNext bool
}

func (p *pager) Next() bool {
	return p.hasNext
}

func (p *pager) Scan(set *host.HostSet) error {
	hs, err := p.op.Query()
	if err != nil {
		return err
	}
	if hs.Length() > 0 {
		p.hasNext = true
	}
}
