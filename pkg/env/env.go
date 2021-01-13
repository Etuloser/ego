package env

import (
	"os"
	"path/filepath"
)

// GetBaseDir Comment
func GetBaseDir() string {
	basedir, _ := os.Getwd()
	return basedir
}

// GetEnvPath Comment
func GetEnvPath() string {
	basedir := GetBaseDir()
	envPath := filepath.Join(basedir, ".env")
	return envPath
}
