package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(addr string, handler *gin.Engine) server {
	s := endless.NewServer(addr, handler)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	return s
}
