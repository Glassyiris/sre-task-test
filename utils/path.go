package utils

import (
	"os"
)

func InitPath() {
	path := "./avatar"
	os.MkdirAll(path, os.ModePerm)
}
