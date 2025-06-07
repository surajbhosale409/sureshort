package service

func (s *Service) registerRoutes() {
	// app routes
	appRoutes := s.echoService.Group("/app")
	// shorten url routes
	appRoutes.GET("/create", s.shortenURLHandler)
	appRoutes.POST("/create", s.shortenURLHandler)
	// metrics route
	appRoutes.GET("/metrics", s.metricsHandler)

	// redirect route
	s.echoService.GET("/:url", s.redirectHanlder)
}
