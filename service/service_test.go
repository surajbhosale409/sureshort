package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	t.Run("creates new service with default config", func(t *testing.T) {
		service := NewService(&Config{})
		assert.NotNil(t, service)
		assert.Equal(t, service.config, &Config{
			ServiceName: defaultServiceName,
			Address:     defaultListenerAddress,
			Port:        defaultPort,
		})
		assert.Equal(t, len(service.echoService.Routes()), 4)
	})

	t.Run("creates new service with custom config", func(t *testing.T) {
		config := &Config{
			ServiceName: "test-service",
			Address:     "localhost",
			Port:        "443",
		}
		service := NewService(config)
		assert.NotNil(t, service)
		assert.Equal(t, service.config, config)
		assert.Equal(t, len(service.echoService.Routes()), 4)
	})
}

// func TestShortner_OriginalURL(t *testing.T) {

// 	// expecting short versions to follow crc32 hash
// 	testData := []struct {
// 		expectedOriginalURL string
// 		shortenedURL        string
// 	}{
// 		{"google.com", "e14f0993"},
// 		{"amazon.com", "2060fdd1"},
// 		{"facebook.com", "22ae9a07"},
// 	}

// 	for _, test := range testData {
// 		t.Run(fmt.Sprintf("test OriginalURL - %s", test.shortenedURL), func(t *testing.T) {
// 			ShortenURL(test.expectedOriginalURL)

// 			if resultURL, err := shortner.OriginalURL(test.shortenedURL); err != nil || resultURL != test.expectedOriginalURL {
// 				t.Errorf("resultant original url is not matching the expected one, err=%s, shortenedURl=%s, expectedOriginalURL=%s, got=%s",
// 					err, test.shortenedURL, test.expectedOriginalURL, resultURL)
// 			}
// 		})
// 	}

// }
