package server

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

type Options struct {
	Host    string
	Port    string
	LogPath string
}
type Server struct {
	app     *router.Router
	Options *Options
	Logger  zerolog.Logger
}

func New(options ...*Options) (s Server) {
	for _, opt := range options {
		if opt != nil {
			defaultOptions = opt
		}
	}
	s.app = router.New()
	s.Logger = Logger(s.Options.LogPath)
	return
}

var defaultOptions = &Options{
	Host:    "localhost",
	Port:    ":8080",
	LogPath: "../logs",
}

func (s Server) RUN(ports ...string) {
	for _, p := range ports {
		if p != "" {
			defaultOptions.Port = p
		}
	}
	log.Fatal(fasthttp.ListenAndServe(s.Options.Port, s.app.Handler))
}
