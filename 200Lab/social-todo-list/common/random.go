package common

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

// GenerateSalt generates a random salt string of the specified length.
func GenerateSalt(length int) (string, error) {
	// Calculate the number of bytes needed.
	bytesNeeded := length / 4 * 3
	if length%4 > 0 {
		bytesNeeded = (length/4 + 1) * 3
	}

	// Generate random bytes.
	randomBytes := make([]byte, bytesNeeded)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64.
	salt := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim the excess characters to match the desired length.
	salt = salt[:length]

	return salt, nil
}

type md5Hash struct{}

func NewMd5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}