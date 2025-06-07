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

	shortenedURL := ShortenURL(reqParams.URL)
	s.urlStore.Store(shortenedURL, reqParams.URL)
	return c.JSON(http.StatusOK, ShortenURLResponseParams{
		ShortenedURL: shortenedURL,
	})
}

func (s *Service) redirectHanlder(c echo.Context) (err error) {
	shortenedURL := c.Param("shortened_url")

	// Lookup original URL
	originalURL, ok := s.urlStore.Load(shortenedURL)
	if !ok || originalURL == "" {
		return c.String(http.StatusNotFound, "Short URL not found")
	}

	// Redirect to original URL
	return c.Redirect(http.StatusMovedPermanently, originalURL.(string))
}
