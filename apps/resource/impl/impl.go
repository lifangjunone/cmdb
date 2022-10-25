package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/apps/resource"
	"github.com/lifangjunone/cmdb/conf"
	registry "github.com/lifangjunone/cmdb/service_registry"
	"google.golang.org/grpc"
)

var (
	svr = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger
	resource.UnimplementedServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.log = zap.L().Named(s.Name())
	s.db = db
	return nil
}

func (s *service) Name() string {
	return resource.AppName
}

func (s *service) Registry(server *grpc.Server) {
	resource.RegisterServiceServer(server, svr)
}

func init() {
	registry.RegistryGrpcApp(svr)
}
