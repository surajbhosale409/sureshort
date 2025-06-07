package pkg

import (
	"fmt"
	"hash/crc32"
)

// Crc32Hash creates short version of URL using CRC32 hashing
func Crc32Hash(value string) (hash string) {
	checksum := crc32.ChecksumIEEE([]byte(value))
	hash = fmt.Sprintf("%08x", checksum)
	return
}
