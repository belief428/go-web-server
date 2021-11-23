package cmd

import (
	"flag"
	"github.com/belief428/go-web-server/router"
	"github.com/belief428/go-web-server/web"
	figure "github.com/common-nighthawk/go-figure"
	"runtime"
)

var ConfigFile = flag.String("config", "./config/config-default.yaml", "config.yaml")

func NewServer(funcs ...func()) {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	figure.NewFigure("Serve Start", "", true).Print()

	if funcs != nil {
		for _, f := range funcs {
			f()
		}
	}
	web.NewWeb()(&web.Config{
		Addr:           "",
		ReadTimeout:    0,
		WriteTimeout:   0,
		IdleTimeout:    0,
		MaxHeaderBytes: 0,
	}).Start(router.NewRouter(&router.Config{
		Mode:    "",
		IsCors:  false,
		Handler: nil,
	}).Handler)
}
