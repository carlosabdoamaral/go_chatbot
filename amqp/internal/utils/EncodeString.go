package utils

import (
	"encoding/base64"
)

func EncodeString(s string) string {
	return _ToBase64(s)
}

func _ToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// func _ToSha1(data []byte, s string) []byte {
// 	block, _ := aes.NewCipher([]byte(createHash(s)))
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		panic(err.Error())
// 	}
// 	ciphertext := gcm.Seal(nonce, nonce, data, nil)
// 	return ciphertext
// }

// func createHash(key string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(key))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }
