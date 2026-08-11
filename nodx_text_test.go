package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestNodeText(t *testing.T) {
	t.Run("Text render", func(t *testing.T) {
		expected := "Hello, World!"
		node := newNodeText(expected)
		assert.Render(t, expected, node)
	})

	t.Run("Empty text render", func(t *testing.T) {
		expected := ""
		node := newNodeText("")
		assert.Render(t, expected, node)
	})

	t.Run("Not initialized nodeText", func(t *testing.T) {
		expected := ""
		node := nodeText{}
		assert.Render(t, expected, node)
	})
}
