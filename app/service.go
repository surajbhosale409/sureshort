package app

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sync"
)

type ShortnerService struct {
	inMemoryStore sync.Map
}

func NewShortnerService() *ShortnerService {
	return &ShortnerService{
		inMemoryStore: sync.Map{},
	}
}

func (s *ShortnerService) ShortenURL(url string) (shortenedURL string) {
	checksum := crc32.ChecksumIEEE([]byte(url))
	shortenedURL = fmt.Sprintf("%08x", checksum)
	s.inMemoryStore.Store(shortenedURL, url)
	return
}

func (s *ShortnerService) OriginalURL(shortenedURL string) (url string, err error) {
	if val, ok := s.inMemoryStore.Load(shortenedURL); !ok {
		err = errors.New("url not found")
	} else {
		url = val.(string)
	}
	return
}
