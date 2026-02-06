package test

import (
	"fmt"
	"testing"

	tools "github.com/ziv/sonder/tools/nil"
)

func TestIsNil(t *testing.T) {
	var p *int
	var i interface{} = p
	fmt.Printf("(i == nil): %v\n", (i == nil))
	fmt.Printf("tools.IsNil(i): %v\n", tools.IsNil(i))
}
