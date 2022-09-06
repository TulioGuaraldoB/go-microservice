package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tulioguaraldo/go-microservice/core/movie"
)

type ServerConfig struct {
	Port   string
	Server *gin.Engine
}

func NewServer() ServerConfig {
	return ServerConfig{
		Port:   os.Getenv("HTTP_PORT"),
		Server: gin.Default(),
	}
}

func (s *ServerConfig) RunServer() {
	log.Fatal(movie.GetRoutes().Run(":" + s.Port))
}
