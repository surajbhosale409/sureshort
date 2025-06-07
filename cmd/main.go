package main

import "github.com/surajbhosale409/sureshort/service"

func main() {
	config := service.Config{
		ServiceName: "sureshort",
		Port:        "80",
	}

	urlShortnerService := service.NewService(config)
	urlShortnerService.Serve()
}
