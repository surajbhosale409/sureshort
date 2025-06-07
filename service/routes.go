package service

func (s *Service) registerRoutes() {
	// app routes
	appRoutes := s.echoService.Group("/app")
	// shorten url routes
	appRoutes.GET("/create", s.shortenURLHandler)
	appRoutes.POST("/create", s.shortenURLHandler)

	// redirect route
	s.echoService.GET("/:shortened_url", s.redirectHanlder)
}
