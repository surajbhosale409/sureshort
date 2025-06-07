package service

import (
	"fmt"
	"hash/crc32"
)

// ShortenURL creates short version of URL using CRC32 hashing
func ShortenURL(url string) (shortenedURL string) {
	checksum := crc32.ChecksumIEEE([]byte(url))
	shortenedURL = fmt.Sprintf("%08x", checksum)
	return
}
