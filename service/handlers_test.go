package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupService(t *testing.T) *Service {
	service := NewService(Config{})
	assert.NotNil(t, service)
	assert.Equal(t, service.config, Config{
		ServiceName: defaultServiceName,
		Address:     defaultListenerAddress,
		Port:        defaultPort,
	})
	assert.Equal(t, len(service.echoService.Routes()), 3)
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
			ShortenedURL: "http://example.com/e14f0993",
		}

		if assert.NoError(t, service.shortenURLHandler(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			response := ShortenURLResponseParams{}
			json.Unmarshal(rec.Body.Bytes(), &response)
			assert.Equal(t, expectedResponse, response)
		}

	})

}

func TestRedirectHandler(t *testing.T) {
	service := setupService(t)

	e := echo.New()

	// shorten a valid url as test setup
	req := httptest.NewRequest(http.MethodGet, "/create?url=google.com", nil)
	req.Header.Add("Accept", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedResponse := ShortenURLResponseParams{
		ShortenedURL: "http://example.com/e14f0993",
	}

	if assert.NoError(t, service.shortenURLHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		response := ShortenURLResponseParams{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, expectedResponse, response)
	}

	t.Run("redirects to orignal url when valid shortened url is passed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/e14f0993", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/:url")
		c.SetParamNames("url")
		c.SetParamValues("e14f0993")

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
