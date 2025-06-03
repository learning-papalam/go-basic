package bot

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptyEnv(key string) []byte {
	block, _ := aes.NewCipher([]byte(key))
	aesGCM, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesGCM.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	return aesGCM.Seal(nonce, nonce, []byte("password123"), nil)
}
