package go_web_server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Web struct {
	httpServer *http.Server
	*Config
}

type Config struct {
	Addr           string
	ReadTimeout    int
	WriteTimeout   int
	IdleTimeout    int
	MaxHeaderBytes int
}

type WebHandle func(config *Config) *Web

var (
	DefaultConfig = &Config{
		Addr:           "127.0.0.1",
		ReadTimeout:    10,
		WriteTimeout:   10,
		IdleTimeout:    10,
		MaxHeaderBytes: 1 << 20,
	}
)

func (this *Web) stop() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGTERM)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				err = fmt.Errorf("internal error: %v", err)
			}
		}()
		<-c
		defer close(c)
		signal.Stop(c)
		fmt.Printf("Http Server Stop")
		os.Exit(0)
	}()
}

func (this *Web) Start(handler http.HandlerFunc) {
	this.httpServer = &http.Server{
		Addr:           this.Addr,
		Handler:        handler,
		ReadTimeout:    time.Duration(this.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(this.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(this.IdleTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	this.stop()

	fmt.Println("Http Server Start")
	fmt.Printf("Http Server Address - %v\n", this.Addr)

	_ = this.httpServer.ListenAndServe()
}

func NewWeb() WebHandle {
	return func(config *Config) *Web {
		if config == nil {
			config = DefaultConfig
		}
		return &Web{Config: config}
	}
}
