package impl

import (
"database/sql"


	

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/lifangjunone/cmdb/apps/book"
	"github.com/lifangjunone/cmdb/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
db   *sql.DB

	log  logger.Logger
	book.UnimplementedServiceServer
}

func (s *service) Config() error {
db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db


	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Name() string {
	return book.AppName
}

func (s *service) Registry(server *grpc.Server) {
	book.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}