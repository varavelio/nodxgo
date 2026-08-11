package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestNodeElement(t *testing.T) {
	t.Run("Basic element render", func(t *testing.T) {
		tag := "div"
		expected := "<div></div>"
		node := newNodeElement(false, tag)
		assert.Render(t, expected, node)
	})

	t.Run("Basic void element", func(t *testing.T) {
		tag := "img"
		expected := "<img>"
		node := newNodeElement(true, tag)
		assert.Render(t, expected, node)
	})

	t.Run("Element with text", func(t *testing.T) {
		tag := "p"
		children := []Node{
			newNodeText("Hello, World!"),
		}
		expected := "<p>Hello, World!</p>"
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Void element with text", func(t *testing.T) {
		tag := "link"
		children := []Node{
			newNodeText("Hello, World!"),
		}
		expected := "<link>"
		node := newNodeElement(true, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with attributes", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("class", "container"),
			newNodeAttribute("id", "main"),
		}
		expected := `<div class="container" id="main"></div>`
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Void element with attributes", func(t *testing.T) {
		tag := "img"
		children := []Node{
			newNodeAttribute("src", "image.jpg"),
			newNodeAttribute("alt", "Image"),
		}
		expected := `<img src="image.jpg" alt="Image">`
		node := newNodeElement(true, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with text and attributes", func(t *testing.T) {
		tag := "a"
		children := []Node{
			newNodeAttribute("href", "#"),
			newNodeAttribute("class", "btn"),
			newNodeText("Click"),
		}
		expected := `<a href="#" class="btn">Click</a>`
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Void element with text and attributes", func(t *testing.T) {
		tag := "link"
		children := []Node{
			newNodeAttribute("rel", "stylesheet"),
			newNodeAttribute("href", "style.css"),
			newNodeText("This will be ignored"),
		}
		expected := `<link rel="stylesheet" href="style.css">`
		node := newNodeElement(true, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with children in mixed order", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("class", "container"),
			newNodeText("Hello"),
			newNodeAttribute("id", "main"),
			newNodeText(", World!"),
			newNodeAttribute("style", "color: red;"),
			newNodeText(" - nodx"),
		}
		expected := `<div class="container" id="main" style="color: red;">Hello, World! - nodx</div>`
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Void element with children in mixed order", func(t *testing.T) {
		tag := "img"
		children := []Node{
			newNodeAttribute("src", "image.jpg"),
			newNodeText("This will be ignored"),
			newNodeAttribute("alt", "Image"),
			newNodeText("This will also be ignored"),
		}
		expected := `<img src="image.jpg" alt="Image">`
		node := newNodeElement(true, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with nested elements", func(t *testing.T) {
		tag := "ul"
		children := []Node{
			newNodeElement(false, "li", newNodeText("Item 1")),
			newNodeElement(false, "li", newNodeText("Item 2")),
		}
		expected := "<ul><li>Item 1</li><li>Item 2</li></ul>"
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with duplicate attributes", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("class", "container"),
			newNodeAttribute("class", "main"),
		}
		expected := `<div class="container" class="main"></div>`
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with special attributes", func(t *testing.T) {
		tag := "button"
		children := []Node{
			newNodeAttribute("data-toggle", "modal"),
			newNodeAttribute("aria-label", "Open Modal"),
		}
		expected := `<button data-toggle="modal" aria-label="Open Modal"></button>`
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with empty tag", func(t *testing.T) {
		node := newNodeElement(false, "")
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Element with nil child", func(t *testing.T) {
		tag := "div"
		children := []Node{
			nil,
			newNodeText("Content"),
		}
		expected := "<div>Content</div>"
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Complex element with many nested levels", func(t *testing.T) {
		tag := "html"
		children := []Node{
			newNodeElement(false, "head",
				newNodeElement(false, "title", newNodeText("Hello, World!")),
				newNodeElement(true, "meta", newNodeAttribute("charset", "UTF-8")),
			),
			newNodeElement(false, "body",
				newNodeElement(false, "div",
					newNodeAttribute("class", "container"),
					newNodeElement(false, "h1", newNodeText("Hello, World!")),
					newNodeElement(false, "p", newNodeText("This is a paragraph.")),
				),
			),
		}
		expected := "<html><head><title>Hello, World!</title><meta charset=\"UTF-8\"></head><body><div class=\"container\"><h1>Hello, World!</h1><p>This is a paragraph.</p></div></body></html>"
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with mixed text and element children", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeText("Hello, "),
			newNodeElement(false, "strong", newNodeText("World")),
			newNodeText("!"),
		}
		expected := "<div>Hello, <strong>World</strong>!</div>"
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})

	t.Run("Element with xss injection", func(t *testing.T) {
		tag := "div"
		children := []Node{
			newNodeAttribute("style", `alert("Hello, World!");`),
			newNodeText("This is a paragraph."),
		}
		expected := `<div style="alert(&quot;Hello, World!&quot;);">This is a paragraph.</div>`
		node := newNodeElement(false, tag, children...)
		assert.Render(t, expected, node)
	})
}
