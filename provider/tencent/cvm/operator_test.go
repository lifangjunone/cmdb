package cvm

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/host"
	"github.com/lifangjunone/cmdb/provider/tencent/connectivity"
	tx_cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"testing"
)

var (
	op *CVMOperator
)

func TestQuery(t *testing.T) {
	req := tx_cvm.NewDescribeInstancesRequest()
	set, err := op.Query(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestPagerQuery(t *testing.T) {
	pg := newPager(op)
	set := host.NewHostSet()
	for pg.Next() {
		if err := pg.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		t.Log(set)
	}
}

func init() {
	err := connectivity.LoadClientFromEnv()
	id := "xxx"
	key := "xxx"
	region := "xxx"
	connectivity.SetClient(connectivity.NewTencentCloudClient(id, key, region))
	if err != nil {
		panic(err)
	}
	zap.DevelopmentSetup()
	op = NewCVMOperator(connectivity.C().CvmClient())
}
