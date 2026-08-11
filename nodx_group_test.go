package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestNodeGroup(t *testing.T) {
	t.Run("Render empty", func(t *testing.T) {
		node := newNodeGroup()
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Render basic", func(t *testing.T) {
		node := newNodeGroup(newNodeText("Hello, World!"))
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("Render group with multiple text nodes", func(t *testing.T) {
		node := newNodeGroup(
			newNodeText("Hello"),
			newNodeText(", "),
			newNodeText("World!"),
		)
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("Render group with mixed nodes", func(t *testing.T) {
		node := newNodeGroup(
			newNodeText("Hello, "),
			newNodeElement(false, "strong", newNodeText("World")),
			newNodeText("!"),
		)
		expected := "Hello, <strong>World</strong>!"
		assert.Render(t, expected, node)
	})

	t.Run("Render group with nil node", func(t *testing.T) {
		node := newNodeGroup(
			newNodeText("Hello"),
			nil,
			newNodeText(", World!"),
		)
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("RenderString for group", func(t *testing.T) {
		node := newNodeGroup(
			newNodeText("Hello"),
			newNodeText(", "),
			newNodeText("World!"),
		)
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("RenderBytes for group", func(t *testing.T) {
		node := newNodeGroup(
			newNodeText("Hello"),
			newNodeText(", "),
			newNodeText("World!"),
		)
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("Group with nested groups", func(t *testing.T) {
		node := newNodeGroup(
			newNodeGroup(
				newNodeText("Hello, "),
				newNodeElement(false, "strong", newNodeText("World")),
			),
			newNodeText("!"),
		)
		expected := "Hello, <strong>World</strong>!"
		assert.Render(t, expected, node)
	})

	t.Run("Empty group RenderString", func(t *testing.T) {
		node := newNodeGroup()
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Empty group RenderBytes", func(t *testing.T) {
		node := newNodeGroup()
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Expand group as div child element", func(t *testing.T) {
		node := El(
			"div",
			newNodeGroup(
				newNodeText("Hello, World!"),
			),
		)
		expected := "<div>Hello, World!</div>"
		assert.Render(t, expected, node)
	})

	t.Run("Expand group as div child attribute", func(t *testing.T) {
		node := El(
			"div",
			newNodeGroup(
				Attr("class", "test"),
			),
			newNodeText("Hello, World!"),
		)
		expected := `<div class="test">Hello, World!</div>`
		assert.Render(t, expected, node)
	})

	t.Run("Expand group as div child element and attribute", func(t *testing.T) {
		node := El(
			"div",
			newNodeGroup(
				Attr("class", "test"),
				newNodeText("Hello, World!"),
			),
		)
		expected := `<div class="test">Hello, World!</div>`
		assert.Render(t, expected, node)
	})

	t.Run("Expand complex nested group", func(t *testing.T) {
		node := Div(
			Class("test1"),
			Group(
				Div(
					Class("test2"),
					Group(
						SpanEl(Text("Hello, ")),
						Group(
							SpanEl(Text("World!")),
						),
					),
					Group(Group(Group(Group(Group(Group(Group(Group(Group(
						SpanEl(Text("NodX")),
					))))))))),
				),
			),
			SpanEl(Text("Hello, World!")),
		)
		expected := `<div class="test1"><div class="test2"><span>Hello, </span><span>World!</span><span>NodX</span></div><span>Hello, World!</span></div>`
		assert.Render(t, expected, node)
	})
}
