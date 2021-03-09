package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func GetName(name string) string {
	sp := strings.Split(name, ".")
	n := len(sp)
	t := sp[n-1]
	h := md5.New()
	h.Write([]byte(name))
	return hex.EncodeToString(h.Sum(nil))  + "." + t
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}