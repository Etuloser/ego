package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config, err := NewConfig()
	fmt.Println(config, err)
}
