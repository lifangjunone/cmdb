package service_registry

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful/v3"
)

var (
	restfulApps = map[string]RESTFulApp{}
)

// RESTFulApp Http服务的实例
type RESTFulApp interface {
	Registry(*restful.WebService)
	Config() error
	Name() string
	Version() string
}

// RegistryRESTFulApp 服务实例注册
func RegistryRESTFulApp(app RESTFulApp) {
	// 已经注册的服务禁止再次注册
	_, ok := restfulApps[app.Name()]
	if ok {
		panic(fmt.Sprintf("http app %s has registed", app.Name()))
	}

	restfulApps[app.Name()] = app
}

// LoadedRESTFulApp 查询加载成功的服务
func LoadedRESTFulApp() (apps []string) {
	for k := range restfulApps {
		apps = append(apps, k)
	}
	return
}

func GetRESTFulApp(name string) RESTFulApp {
	app, ok := restfulApps[name]
	if !ok {
		panic(fmt.Sprintf("http app %s not registed", name))
	}

	return app
}

// LoadRESTFulApp 装载所有的http app
func LoadRESTFulApp(pathPrefix string, root *restful.Container) {
	for _, api := range restfulApps {
		pathPrefix = strings.TrimSuffix(pathPrefix, "/")
		ws := new(restful.WebService)
		ws.
			Path(fmt.Sprintf("%s/%s/%s", pathPrefix, api.Version(), api.Name())).
			Consumes(restful.MIME_JSON, restful.MIME_XML).
			Produces(restful.MIME_JSON, restful.MIME_XML)

		api.Registry(ws)
		root.Add(ws)
	}
}
