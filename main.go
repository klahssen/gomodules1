package main

import (
	"fmt"

	"github.com/klahssen/gomodules1/internal/pkg1"
)

const (
	version = "v0.1.0"
)

func main() {
	fmt.Printf("gomodule version '%s'\n", version)
	pkg1.Greet()
}
