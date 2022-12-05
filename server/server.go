package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Root        string       `yaml:"root"`
	Port        int          `yaml:"port"`
	Secure      bool         `yaml:"secure"`
	Certificate *Certificate `yaml:"certificate"`
}

type Certificate struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

func (s *Server) Init() error {

	app := fiber.New(fiber.Config{
		AppName: "SwaggerUI",
	})

	app.Static("/", s.Root)
	//app.Get("/")

	if s.Secure && s.Certificate != nil {
		return app.ListenTLS(fmt.Sprintf(":%d", s.Port), s.Certificate.Cert, s.Certificate.Key)
	}

	return app.Listen(fmt.Sprintf(":%d", s.Port))
}
