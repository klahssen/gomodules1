package pkg1

import "fmt"

const (
	version = "v0.0.2"
)

//Greet method
func Greet() {
	fmt.Printf("Greet version '%s'\n", version)
}
