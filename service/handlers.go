package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/surajbhosale409/sureshort/pkg"
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

	if !hasScheme(reqParams.URL) {
		reqParams.URL = fmt.Sprintf("http://%s", reqParams.URL)
	}

	shortenedURL := pkg.Crc32Hash(reqParams.URL)
	s.urlStore.Store(shortenedURL, reqParams.URL)

	// asynchronously record stats
	go s.recordStats(reqParams.URL)

	shortenedURL = fmt.Sprintf("http://%s/%s", c.Request().Host, shortenedURL)
	acceptHeader := c.Request().Header.Get("Accept")
	switch acceptHeader {
	case "application/json":
		return c.JSON(http.StatusOK, ShortenURLResponseParams{
			ShortenedURL: shortenedURL,
		})
	default: //return text/html response by default
		return c.HTML(http.StatusOK, fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", shortenedURL, shortenedURL))
	}
}

func hasScheme(rawurl string) bool {
	u, err := url.Parse(rawurl)
	if err != nil {
		return false // invalid URL, treat as no scheme
	}
	return u.Scheme != ""
}

func (s *Service) redirectHanlder(c echo.Context) (err error) {
	shortenedURL := c.Param("url")

	// Lookup original URL
	url, ok := s.urlStore.Load(shortenedURL)
	if !ok || url == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "shortened url not found")
	}

	originalURL := url.(string)

	// Redirect to original URL
	c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("Expires", "0")
	c.Response().Header().Set("Surrogate-Control", "no-store")
	return c.Redirect(http.StatusFound, originalURL)
}

func (s *Service) metricsHandler(c echo.Context) (err error) {
	response := strings.Join(s.stats.Top(3), "\n")
	return c.String(http.StatusOK, response)
}
