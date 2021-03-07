package utils

import (
	"fmt"
	"testing"
)

func TestInitDB(t *testing.T) {
	InitDB()
	fmt.Println(Configs.Database.Dsn)
}
