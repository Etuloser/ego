package test

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	fmt.Println(os.Getwd())
}
