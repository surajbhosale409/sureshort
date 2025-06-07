package service

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	ServiceName string `json:"serviceName" yaml:"serviceName"`
	Address     string `json:"address" yaml:"serviceName"`
	Port        string `json:"port" yaml:"serviceName"`
}

type Service struct {
	config      Config
	urlStore    sync.Map
	echoService *echo.Echo
}

const (
	defaultServiceName     = "sureshort"
	defaultListenerAddress = "*"
	defaultPort            = "80"
)

func NewService(config Config) (service *Service) {

	// if address is not provided, set * to listen on addresses of all interfaces
	if config.Address == "" {
		config.Address = defaultListenerAddress
	}
	// if port is not provided, use default port 80
	if config.Port == "" {
		config.Port = defaultPort
	}

	// use default servicename if not provided
	if config.ServiceName == "" {
		config.ServiceName = defaultServiceName
	}

	service = &Service{
		config: config,
	}
	service.Initialise()

	return
}

func (s *Service) Initialise() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	s.echoService = e
	s.registerRoutes()
}

func (s *Service) Serve() {
	s.echoService.Logger.Fatal(s.echoService.Start(s.config.Address + ":" + s.config.Port))
}

func (s *Service) recordStats(url string) {

}
