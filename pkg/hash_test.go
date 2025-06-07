package pkg

import (
	"fmt"
	"testing"
)

func TestHash_Crc32Hash(t *testing.T) {

	// expecting short versions to follow crc32 hash
	testData := []struct {
		url               string
		expectedCrc32Hash string
	}{
		{"google.com", "e14f0993"},
		{"amazon.com", "2060fdd1"},
		{"facebook.com", "22ae9a07"},
	}

	for _, test := range testData {
		t.Run(fmt.Sprintf("test Crc32Hash - %s", test.url), func(t *testing.T) {
			if resultURL := Crc32Hash(test.url); resultURL != test.expectedCrc32Hash {
				t.Errorf("shortened url is not following crc32 hash, originalURL=%s, expectedCrc32Hash=%s, got=%s",
					test.url, test.expectedCrc32Hash, resultURL)
			}
		})
	}

}
