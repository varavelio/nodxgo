package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestNodeAttribute(t *testing.T) {
	t.Run("Attribute render", func(t *testing.T) {
		key := "class"
		value := "container"
		expected := `class="container"`
		node := newNodeAttribute(key, value)
		assert.Render(t, expected, node)
	})

	t.Run("Empty value attribute render", func(t *testing.T) {
		key := "class"
		value := ""
		expected := `class=""`
		node := newNodeAttribute(key, value)
		assert.Render(t, expected, node)
	})

	t.Run("Empty key attribute render", func(t *testing.T) {
		key := ""
		value := "container"
		expected := ""
		node := newNodeAttribute(key, value)
		assert.Render(t, expected, node)
	})

	t.Run("Empty key and value attribute render", func(t *testing.T) {
		key := ""
		value := ""
		expected := ""
		node := newNodeAttribute(key, value)
		assert.Render(t, expected, node)
	})

	t.Run("Double quote in value attribute render", func(t *testing.T) {
		key := "class"
		value := `container"container'container`
		expected := `class="container&quot;container&#39;container"`
		node := newNodeAttribute(key, value)
		assert.Render(t, expected, node)
	})

	t.Run("Attribute with no value at all", func(t *testing.T) {
		key := "disabled"
		expected := `disabled`
		node := newNodeAttribute(key)
		assert.Render(t, expected, node)
	})
}
