package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

// encriptamos password
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))                // encriptamos
	return hex.EncodeToString(hash.Sum(nil)) // usamos hex para retornar un string
}
