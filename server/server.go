package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Name        string       `json:"name"`
	Root        string       `yaml:"root"`
	ApiUri      string       `yaml:"apiUri"`
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
		AppName: s.Name,
	})

	app.Static("/", s.Root)
	app.Use(func(c *fiber.Ctx) error {
		return proxy.Do(c, s.ApiUri+c.Path())
	})

	if s.Secure && s.Certificate != nil {
		return app.ListenTLS(fmt.Sprintf(":%d", s.Port), s.Certificate.Cert, s.Certificate.Key)
	}

	return app.Listen(fmt.Sprintf(":%d", s.Port))
}
