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

	cases := []struct {
		name string
		make func() fyne.CanvasObject
	}{
		{"Colors", makeColors},
		{"Icons", makeIcons},
		{"Sizes", makeSizes},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			obj := c.make()
			assert.NotNil(t, obj)
		})
	}
}
