package secret_test

import (
	"context"
	"github.com/infraboard/mcube/logger/zap"
	_ "github.com/lifangjunone/cmdb/apps/all"
	"github.com/lifangjunone/cmdb/apps/secret"
	"github.com/lifangjunone/cmdb/conf"
	"github.com/lifangjunone/cmdb/service_registry"
	"testing"
)

var (
	ins secret.ServiceServer
)

func TestDescribeSecret(t *testing.T) {
	ss, err := ins.QuerySecret(context.Background(), secret.NewQuerySecretRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func init() {
	err := conf.LoadConfigFromToml("../../etc/config.toml")
	if err != nil {
		panic(err)
	}
	zap.DevelopmentSetup()
	service_registry.InitAllApp()
	ins = service_registry.GetGrpcApp(secret.AppName).(secret.ServiceServer)

}
