package cvm

import (
	"context"
	"github.com/infraboard/mcube/flowcontrol/tokenbucket"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/host"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	tx_cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"time"
)

type pager struct {
	req        *cvm.DescribeInstancesRequest
	op         *CVMOperator
	hasNext    bool
	pageNumber int64
	pageSize   int64
	log        logger.Logger
	tb         *tokenbucket.Bucket
}

func newPager(op *CVMOperator) host.Pager {
	p := &pager{
		op:         op,
		hasNext:    true,
		pageNumber: 1,
		pageSize:   20,
		log:        zap.L().Named("CVM"),
		tb:         tokenbucket.NewBucket(3*time.Second, 3),
	}
	p.req = tx_cvm.NewDescribeInstancesRequest()
	p.req.Limit = &p.pageSize
	p.req.Offset = p.offset()
	return p
}

func NewPagger(rate float64, op *CVMOperator) host.Pager {
	p := &pager{
		op:         op,
		hasNext:    true,
		pageNumber: 1,
		pageSize:   20,
		log:        zap.L().Named("CVM"),
		tb:         tokenbucket.NewBucketWithRate(rate, 3),
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
	p.tb.Wait(1)
	p.req.Offset = p.offset()
	p.req.Limit = &p.pageSize
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
	*set = *hs.Clone()
	p.pageNumber++
	return nil
}
