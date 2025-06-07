package main

import (
	"fmt"

	"github.com/surajbhosale409/sureshort/app"
)

func main() {

	service := app.NewShortner()
	googleShort := service.ShortenURL("google.com")
	fmt.Println("facebook.com", service.ShortenURL("facebook.com"))
	fmt.Println("amazon.com", service.ShortenURL("amazon.com"))
	fmt.Println("google.com", service.ShortenURL("google.com"))

	if url, err := service.OriginalURL(googleShort); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(googleShort, url)
	}

}
