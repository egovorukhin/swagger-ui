package main

import (
	"fmt"
	info "github.com/egovorukhin/egoappinfo"
	"github.com/egovorukhin/egoconf"
	"log"
	"os"
	"os/signal"
	"swagger-ui/server"
	"syscall"
)

type Config struct {
	server.Server `yaml:",inline"`
}

func init() {
	info.SetName("Swagger UI")
	info.SetVersion(0, 0, 1)
}

func main() {

	// Канал для получения ошибки, если таковая будет
	errChan := make(chan error, 2)

	// Ждем сигнал от ОС
	go waitSignal(errChan)
	// Стартовая горутина
	go start(errChan)

	err := <-errChan
	if err != nil {
		log.Println(err)
	}
}

func start(errChan chan error) {

	// Загружаем конфигурацию приложения
	cfg := &Config{}
	err := egoconf.Load("config.yml", cfg)
	if err != nil {
		errChan <- err
		return
	}

	// Запускаем сервер
	errChan <- cfg.Server.Init()
}

func waitSignal(errChan chan error) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL) //nolint:govet
	errChan <- fmt.Errorf("%s", <-c)
}
