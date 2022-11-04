package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/host"
	"github.com/lifangjunone/cmdb/apps/secret"
	"github.com/lifangjunone/cmdb/conf"
	"github.com/lifangjunone/cmdb/service_registry"
	"google.golang.org/grpc"
)

var (
	svc = &impl{}
)

type impl struct {
	db   *sql.DB
	log  logger.Logger
	host host.ServiceServer
	secret.UnimplementedServiceServer
}

func (i *impl) Name() string {
	return secret.AppName
}

func (i *impl) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	i.db = db
	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Registry(server *grpc.Server) {
	secret.RegisterServiceServer(server, svc)
}

func init() {
	service_registry.RegistryGrpcApp(svc)
}
