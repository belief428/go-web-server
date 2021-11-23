package router

import (
	"github.com/belief428/go-web-server/utils"
	"net/http"
)

type Handler struct {
	Writer      http.ResponseWriter
	Request     *http.Request
	SkipperUrl  []SkipperURL
	HandlerPass func() // 通过执行的方法
	HandlerFail func() // 未通过执行的方法
}

type SkipperURL func(w http.ResponseWriter, r *http.Request) bool

func AddSkipperURL(url ...string) SkipperURL {
	return func(w http.ResponseWriter, r *http.Request) bool {
		path := r.URL.Path
		return utils.InArray(path, url)
	}
}

func HandleSkipper(handler Handler) {
	if len(handler.SkipperUrl) > 0 && handler.SkipperUrl[0](handler.Writer, handler.Request) {
		if handler.HandlerPass != nil {
			handler.HandlerPass()
		}
		return
	}
	if handler.HandlerFail != nil {
		handler.HandlerFail()
	}
	return
}
