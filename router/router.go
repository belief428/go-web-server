package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler *gin.Engine
}

type Config struct {
	Mode    string
	IsCors  bool
	Handler *gin.Engine
}

func (this *Router) basicApi() {

}

func NewRouter(config *Config) *Router {
	gin.SetMode(config.Mode)

	app := gin.New()

	if config.IsCors {
		//app.Use(Cors())
	}
	//app.NoRoute(NoRouteHandler())
	//app.NoMethod(NoMethodHandler())
	//app.Use(LoggerHandle("log/gin.log", 3))
	//app.Use(TimeoutHandle(time.Second * 30))

	out := &Router{Handler: app}
	out.basicApi()

	return out
}
