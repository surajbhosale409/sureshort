package service

import (
	"fmt"
	"testing"
)

func TestShortner_ShortenURL(t *testing.T) {

	// expecting short versions to follow crc32 hash
	testData := []struct {
		url                  string
		expectedShortenedURL string
	}{
		{"google.com", "e14f0993"},
		{"amazon.com", "2060fdd1"},
		{"facebook.com", "22ae9a07"},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("test ShortenURL - %s", test.url), func(t *testing.T) {
			if resultURL := ShortenURL(test.url); resultURL != test.expectedShortenedURL {
				t.Errorf("shortened url is not following crc32 hash, originalURL=%s, expectedShortenedURL=%s, got=%s",
					test.url, test.expectedShortenedURL, resultURL)
			}
		})
	}

}
