package service

import (
	"net/url"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/surajbhosale409/sureshort/pkg"
)

type Config struct {
	ServiceName string `json:"serviceName" yaml:"serviceName"`
	Address     string `json:"address" yaml:"address"`
	Port        string `json:"port" yaml:"port"`
}

type Service struct {
	config      *Config
	urlStore    sync.Map
	stats       *pkg.Stats
	echoService *echo.Echo
}

const (
	defaultServiceName     = "sureshort"
	defaultListenerAddress = "0.0.0.0"
	defaultPort            = "8080"
)

func NewService(config *Config) (service *Service) {

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
		stats:  pkg.NewStats(),
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

func (s *Service) recordStats(targetURL string) {
	u, err := url.Parse(targetURL)
	if err != nil || u.Host == "" {
		return
	}
	s.stats.Observe(u.Host)
}
