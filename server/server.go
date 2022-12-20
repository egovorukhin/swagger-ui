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
	Timeout Timeout `yaml:"timeout"`
}

type Certificate struct {
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

type Timeout struct {
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
	Idle  int `yaml:"idle"`
}

func (s *Server) Init() error {

	app := fiber.New(fiber.Config{
		AppName: s.Name,
		ReadTimeout:  time.Duration(s.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(s.Timeout.Write) * time.Second,
		IdleTimeout:  time.Duration(s.Timeout.Idle) * time.Second,
	})

	app.Static("/", s.Root)
	app.Use(func(c *fiber.Ctx) error {
		return proxy.Do(c, s.ApiUri+c.Path())
	})
	
	addr := fmt.Sprintf(":%d", s.Port)

	if s.Secure && s.Certificate != nil {
		return app.ListenTLS(addr, s.Certificate.Cert, s.Certificate.Key)
	}

	return app.Listen(addr)
}
