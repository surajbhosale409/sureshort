package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShortenURLRequestParams struct {
	URL string `json:"url" query:"url"`
}

type ShortenURLResponseParams struct {
	ShortenedURL string `json:"shortened_url"`
}

func (s *Service) shortenURLHandler(c echo.Context) (err error) {
	var reqParams ShortenURLRequestParams

	if err := c.Bind(&reqParams); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body")
	}

	if reqParams.URL == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "url cannot be empty")
	}

	// asynchronously record shortening stats
	go s.recordStats(reqParams.URL)

	shortenedURLResp := ShortenURLResponseParams{}
	shortenedURLResp.ShortenedURL = ShortenURL(reqParams.URL)
	return c.JSON(http.StatusOK, shortenedURLResp)
}
