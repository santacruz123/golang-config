package aaa_test

import (
	"goconf/pkg/aaa"
	"os"
	"testing"

	_ "goconf/pkg/utils"

	_ "github.com/joho/godotenv/autoload"
)

func TestWorks(t *testing.T) {
	v := aaa.Print()
	dir, _ := os.Getwd()
	t.Log(v, dir)
}
