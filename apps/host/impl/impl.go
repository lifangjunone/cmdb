package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/lifangjunone/cmdb/service_registry"
	"google.golang.org/grpc"

	"github.com/lifangjunone/cmdb/apps/host"
	"github.com/lifangjunone/cmdb/conf"
)

var (
	svc = &Service{}
)

type Service struct {
	db  *sql.DB
	log logger.Logger
	host.UnimplementedServiceServer
}

func (s *Service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.log = zap.L().Named(s.Name())
	s.db = db
	return nil
}

func (s *Service) Name() string {
	return host.AppName
}

func (s *Service) Registry(server *grpc.Server) {
	host.RegisterServiceServer(server, svc)
}

func init() {
	service_registry.RegistryGrpcApp(svc)
}
