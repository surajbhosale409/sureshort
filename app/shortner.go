package app

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sync"
)

type Shortner struct {
	inMemoryStore sync.Map
}

func NewShortner() *Shortner {
	return &Shortner{
		inMemoryStore: sync.Map{},
	}
}

// ShortenURL creates short version of URL using CRC32 hashing
func (s *Shortner) ShortenURL(url string) (shortenedURL string) {
	checksum := crc32.ChecksumIEEE([]byte(url))
	shortenedURL = fmt.Sprintf("%08x", checksum)
	s.inMemoryStore.Store(shortenedURL, url)
	return
}

// OriginalURL looks-up storage for original URL associated with short version and returns it if found
func (s *Shortner) OriginalURL(shortenedURL string) (url string, err error) {
	if val, ok := s.inMemoryStore.Load(shortenedURL); !ok {
		err = errors.New("url not found")
	} else {
		url = val.(string)
	}
	return
}
