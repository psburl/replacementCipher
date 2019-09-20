package cipher

import (
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// ReplacementCipher its a structure that implements cipher interface
type ReplacementCipher struct {
}

var alphabet = ("abcdefghijklmnopqrstuvwxyz")

// GenerateKey builds a valid key to make replacement cipher
func (ReplacementCipher) GenerateKey(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune(alphabet)
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// Encode method encodes a given text using replacement cipher based on a given key
func (ReplacementCipher) Encode(key, text string) string {
	text = buildTextToAlphabet(text)
	encodedText := ""
	for i, char := range text {
		keyChar := key[i%len(text)]
		sum := strings.Index(alphabet, string(char)) + 1 + strings.Index(alphabet, string(keyChar)) + 1
		encodedText += string(alphabet[sum-1%len(alphabet)])
	}
	return encodedText
}

// Decode method decodes a given text using replacement cipher based on a given key
func (ReplacementCipher) Decode(key, text string) string {
	text = buildTextToAlphabet(text)
	encodedText := ""
	for i, char := range text {
		keyChar := key[i%len(text)]
		sum := (strings.Index(alphabet, string(char)) + 1) - (strings.Index(alphabet, string(keyChar)) + 1)
		if sum < 0 {
			sum += len(alphabet)
		}
		encodedText += string(alphabet[sum-1%len(alphabet)])
	}
	return encodedText
}

func buildTextToAlphabet(text string) string {
	text = strings.ToLower(text)
	reg, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(text, "")
}
