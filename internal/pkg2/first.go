package pkg2

import (
	"testing"

	"github.com/klahssen/tester"
)

func TestOne(t *testing.T) {
	num := 1
	te := tester.NewT(t)
	te.DeepEqual(1, "res", 1, num)
}
