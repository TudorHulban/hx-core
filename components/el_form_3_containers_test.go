package hxcomponents

import (
	"fmt"
	"os"
	"testing"
)

func TestNewForm3Containers(t *testing.T) {
	n := NewFormThreeContainers(
		&ParamsNewFormThreeContainers{},
	)

	fmt.Println(
		n.Render(os.Stdout),
	)
}
