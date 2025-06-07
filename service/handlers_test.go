package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupService(t *testing.T) *Service {
	service := NewService(&Config{})
	assert.NotNil(t, service)
	assert.Equal(t, service.config, &Config{
		ServiceName: defaultServiceName,
		Address:     defaultListenerAddress,
		Port:        defaultPort,
	})
	assert.Equal(t, len(service.echoService.Routes()), 4)
	return service
}

func TestShortenURLHanlder(t *testing.T) {
	service := setupService(t)

	t.Run("fails to shorten url when empty url is passed", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/create", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.Error(t, service.shortenURLHandler(c))

	})

	t.Run("shortens url when valid url is passed", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/create?url=google.com", nil)
		req.Header.Add("Accept", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		expectedResponse := ShortenURLResponseParams{
			ShortenedURL: "http://example.com/7120cf4d",
		}

		if assert.NoError(t, service.shortenURLHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			response := ShortenURLResponseParams{}
			json.Unmarshal(rec.Body.Bytes(), &response)
			assert.Equal(t, expectedResponse, response)
		}

	})

}

func createShortURL(t *testing.T, service *Service, url string) {
	e := echo.New()

	// shorten a valid url as test setup
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/create?url=%s", url), nil)
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, service.shortenURLHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestRedirectHandler(t *testing.T) {
	service := setupService(t)

	createShortURL(t, service, "google.com")

	t.Run("redirects to orignal url when valid shortened url is passed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/7120cf4d", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:url")
		c.SetParamNames("url")
		c.SetParamValues("7120cf4d")

		err := service.redirectHanlder(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusFound, rec.Code)
		assert.Equal(t, "http://google.com", rec.Header().Get("Location"))
	})

	t.Run("returns error wgeb invalid shortened url is passed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/abcdefgh", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:url")
		c.SetParamNames("url")
		c.SetParamValues("abcdefgh")

		err := service.redirectHanlder(c)
		assert.Error(t, err)
	})
}

func TestMetricsHandler(t *testing.T) {
	service := setupService(t)

	createShortURL(t, service, "google.com")
	createShortURL(t, service, "google.com")
	createShortURL(t, service, "google.com")
	createShortURL(t, service, "google.com")

	createShortURL(t, service, "amazon.com")
	createShortURL(t, service, "amazon.com")
	createShortURL(t, service, "amazon.com")

	createShortURL(t, service, "docker.com")

	createShortURL(t, service, "abc.xyz")
	createShortURL(t, service, "abc.xyz")

	t.Run("returns metrics with top 3 domains", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/app/metrics", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := service.metricsHandler(c)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "google.com: 4\namazon.com: 3\nabc.xyz: 2", rec.Body.String())
	})
}
