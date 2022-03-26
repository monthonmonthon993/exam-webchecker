package server

import (
	"log"

	"exam.com/webchecker/internal/controller"
	"github.com/gin-gonic/gin"
)

type server struct {
	*gin.Engine
}

type Server interface {
	InitRoute(c *controller.Controller)
	Run()
}

func NewServer() Server {

	return &server{gin.Default()}
}

func (s *server) InitRoute(c *controller.Controller) {

	v1 := s.Group("/api/v1")
	{
		v1.POST("/webchecker/websites", c.CheckAvalibleWebsites)
	}

}

func (s *server) Run() {
	log.Println("server start")
	s.Engine.Run()
}
