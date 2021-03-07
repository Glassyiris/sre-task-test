package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetName(name string) string {
	h := md5.New()
	h.Write([]byte(name))
	return hex.EncodeToString(h.Sum(nil))
}
