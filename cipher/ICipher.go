package cipher

// ICipher is
type ICipher interface {
	GenerateKey(length int) string
	Encode(string, string) string
	Decode(string, string) string
}
