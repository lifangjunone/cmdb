package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/host"
	"github.com/lifangjunone/cmdb/apps/secret"
	"github.com/lifangjunone/cmdb/apps/task"
	"github.com/lifangjunone/cmdb/conf"
	"github.com/lifangjunone/cmdb/service_registry"
	"google.golang.org/grpc"
)

var (
	svr = &impl{}
)

type impl struct {
	db  *sql.DB
	log logger.Logger
	task.UnimplementedServiceServer

	// 内部	　grpc　调用，直接通过 Server　来调用，即直接调用本地类
	// 外部	　grpc　调用，通过　Client 来调用，通过网络远程调用
	secret secret.ServiceServer
	host   host.ServiceServer
}

func (s *impl) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.log = zap.L().Named(s.Name())
	s.db = db
	s.secret = service_registry.GetGrpcApp(s.Name()).(secret.ServiceServer)
	s.host = service_registry.GetGrpcApp(s.Name()).(host.ServiceServer)
	return nil
}

func (s *impl) Name() string {
	return task.AppName
}

func (s *impl) Registry(server *grpc.Server) {
	task.RegisterServiceServer(server, svr)
}

func init() {
	service_registry.RegistryGrpcApp(svr)
}
