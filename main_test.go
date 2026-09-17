package main

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/stretchr/testify/assert"
)

// TestMakeFunctions is a smoke test which ensures every tab constructor
// runs without panicking and returns a usable canvas object.
func TestMakeFunctions(t *testing.T) {
	testApp := test.NewApp()
	defer testApp.Quit()

	noopCopy := func(string) {}

	cases := []struct {
		name string
		make func(func(string)) fyne.CanvasObject
	}{
		{"Colors", makeColors},
		{"Icons", makeIcons},
		{"Sizes", makeSizes},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			obj := c.make(noopCopy)
			assert.NotNil(t, obj)
		})
	}
}
