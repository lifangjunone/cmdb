package cvm

import (
	"context"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	tx_cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type pager struct {
	req        *cvm.DescribeInstancesRequest
	op         *CVMOperator
	hasNext    bool
	pageNumber int64
	pageSize   int64
	log        logger.Logger
}

func newPager(op *CVMOperator) host.Pager {
	p := &pager{
		op:         op,
		hasNext:    true,
		pageNumber: 1,
		pageSize:   20,
		log:        zap.L().Named("CVM"),
	}
	p.req = tx_cvm.NewDescribeInstancesRequest()
	p.req.Limit = &p.pageSize
	p.req.Offset = p.offset()
	return p
}

func (p *pager) offset() *int64 {
	os := (p.pageNumber - 1) * p.pageSize
	return &os
}

func (p *pager) nextReq() *cvm.DescribeInstancesRequest {
	p.req.Offset = p.offset()
	return p.req
}

func (p *pager) Next() bool {
	return p.hasNext
}

func (p *pager) SetPageSize(ps int64) {
	p.pageSize = ps
}

func (p *pager) Scan(ctx context.Context, set *host.HostSet) error {
	p.log.Debugf("query page: %d", p.pageNumber)
	hs, err := p.op.Query(ctx, p.nextReq())
	if err != nil {
		return err
	}
	if hs.Length() < p.pageSize {
		p.hasNext = false
	}
	p.pageNumber++
	return nil
}