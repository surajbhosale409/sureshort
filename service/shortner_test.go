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

func TestShortner_OriginalURL(t *testing.T) {

	shortner := NewShortner()

	// expecting short versions to follow crc32 hash
	testData := []struct {
		expectedOriginalURL string
		shortenedURL        string
	}{
		{"google.com", "e14f0993"},
		{"amazon.com", "2060fdd1"},
		{"facebook.com", "22ae9a07"},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("test OriginalURL - %s", test.shortenedURL), func(t *testing.T) {
			ShortenURL(test.expectedOriginalURL)

			if resultURL, err := shortner.OriginalURL(test.shortenedURL); err != nil || resultURL != test.expectedOriginalURL {
				t.Errorf("resultant original url is not matching the expected one, err=%s, shortenedURl=%s, expectedOriginalURL=%s, got=%s",
					err, test.shortenedURL, test.expectedOriginalURL, resultURL)
			}
		})
	}

}
