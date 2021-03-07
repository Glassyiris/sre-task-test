package config

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	c := ParseConfig()

	fmt.Println(c.Database.Dsn)
}
