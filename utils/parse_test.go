package utils

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	c := ParseConfig("../config/config.yaml")
	fmt.Println(c.Database.Dsn)
}
