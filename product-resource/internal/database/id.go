package database

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"strings"
)

// NewID creates a new database ID.
func NewID(prefix string) string {
	underscore := ""
	if !strings.Contains(prefix, "_") {
		underscore = "_"
	}

	return prefix + underscore + randomString(42)
}

// defaultLetters is a list of default letters that can be used to make a random
// string when calling String function with no letters provided
var defaultLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// randomString generates a random string according to the size entered and the letters.
// Examples:
// RandString(5) -> "EfPEK"
func randomString(size int, letters ...string) string {
	letterRunes := defaultLetters

	if len(letters) > 0 {
		letterRunes = []rune(letters[0])
	}

	var bb bytes.Buffer
	bb.Grow(size)
	l := uint32(len(letterRunes))

	// on each loop, generate one random rune and append to output
	for i := 0; i < size; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(generateBytes(4))%l])
	}

	return bb.String()
}

// generateBytes generates n random bytes
func generateBytes(n int) []byte {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}
