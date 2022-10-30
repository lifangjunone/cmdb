package service_registry

import (
	"fmt"

	"google.golang.org/grpc"
)

var (
	grpcApps = map[string]GRPCApp{}
)

// GRPCApp GRPC 服务注册实例
type GRPCApp interface {
	Registry(*grpc.Server)
	Config() error
	Name() string
}

// RegistryGrpcApp 服务实例注册
func RegistryGrpcApp(app GRPCApp) {
	_, ok := grpcApps[app.Name()]
	if ok {
		fmt.Sprintf("grpc app %s has registed", app.Name())
	}
	grpcApps[app.Name()] = app
}

// LoadedGrpcApp 查询已注册的服务
func LoadedGrpcApp() (apps []string) {
	for k := range grpcApps {
		apps = append(apps, k)
	}
	return
}

// GetGrpcApp 通过服务名获取对应的服务实例
func GetGrpcApp(name string) GRPCApp {
	app, ok := grpcApps[name]
	if !ok {
		fmt.Sprintf("grpc app %s not registry", name)
		return nil
	}
	return app
}

// LoadGrpcApp 注册所有的服务到　GRPC服务中
func LoadGrpcApp(server *grpc.Server) {
	for _, app := range grpcApps {
		app.Registry(server)
	}
}
