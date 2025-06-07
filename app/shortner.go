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

func (s *Shortner) ShortenURL(url string) (shortenedURL string) {
	checksum := crc32.ChecksumIEEE([]byte(url))
	shortenedURL = fmt.Sprintf("%08x", checksum)
	s.inMemoryStore.Store(shortenedURL, url)
	return
}

func (s *Shortner) OriginalURL(shortenedURL string) (url string, err error) {
	if val, ok := s.inMemoryStore.Load(shortenedURL); !ok {
		err = errors.New("url not found")
	} else {
		url = val.(string)
	}
	return
}
