package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestStyleMap(t *testing.T) {
	t.Run("Render empty StyleMap", func(t *testing.T) {
		node := StyleMap{}
		expected := `style=""`
		assert.Render(t, expected, node)
	})

	t.Run("Render StyleMap with single true style", func(t *testing.T) {
		node := StyleMap{
			"border: 1px solid black": true,
		}
		expected := `style="border: 1px solid black"`
		assert.Render(t, expected, node)
	})

	t.Run("Render StyleMap with multiple true styles", func(t *testing.T) {
		node := StyleMap{
			"border: 1px solid black": true,
			"margin: 10px":            true,
			"color: red":              true,
		}
		expected := `style="border: 1px solid black; color: red; margin: 10px"`
		assert.Render(t, expected, node)
	})

	t.Run("Render StyleMap with mixed true and false styles", func(t *testing.T) {
		node := StyleMap{
			"border: 1px solid black": true,
			"padding: 20px":           false,
			"margin: 10px":            true,
			"color: red":              false,
		}
		expected := `style="border: 1px solid black; margin: 10px"`
		assert.Render(t, expected, node)
	})

	t.Run("Render sorted StyleMap", func(t *testing.T) {
		node := StyleMap{
			"z-index: 1":  true,
			"border: 1px": true,
			"margin: 5px": true,
		}
		expected := `style="border: 1px; margin: 5px; z-index: 1"`
		assert.Render(t, expected, node)
	})

	t.Run("StyleMap with no true values", func(t *testing.T) {
		node := StyleMap{
			"border: 1px solid black": false,
			"margin: 10px":            false,
		}
		expected := `style=""`
		assert.Render(t, expected, node)
	})

	t.Run("StyleMap with special characters in styles", func(t *testing.T) {
		node := StyleMap{
			"content: 'Hello, World!'": true,
			"font-size: 16px":          true,
		}
		expected := `style="content: &#39;Hello, World!&#39;; font-size: 16px"`
		assert.Render(t, expected, node)
	})

	t.Run("StyleMap with conditional styles", func(t *testing.T) {
		showBorder := true
		showPadding := false

		node := StyleMap{
			"border: 1px solid black": showBorder,
			"padding: 20px":           showPadding,
			"margin: 5px":             true,
		}
		expected := `style="border: 1px solid black; margin: 5px"`

		assert.Render(t, expected, node)
	})

	t.Run("Render inside a node", func(t *testing.T) {
		node := El("div", StyleMap{
			"border: 1px solid black": true,
			"margin: 5px":             false,
		})
		expected := `<div style="border: 1px solid black"></div>`

		assert.Render(t, expected, node)
	})
}
