package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestGroup(t *testing.T) {
	t.Run("Group with no nodes", func(t *testing.T) {
		node := Group()
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Group with multiple nodes", func(t *testing.T) {
		node := Group(Text("Hello"), Text("World"))
		expected := "HelloWorld"
		assert.Render(t, expected, node)
	})

	t.Run("Group with some nil nodes", func(t *testing.T) {
		node := Group(Text("Hello"), nil, Text("World"))
		expected := "HelloWorld"
		assert.Render(t, expected, node)
	})

	t.Run("Group with all nil nodes", func(t *testing.T) {
		node := Group(nil, nil, nil)
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Group with mixed children", func(t *testing.T) {
		node := Group(
			El("div",
				Attr("class", "container"),
				Text("Hello"),
				Text("World"),
				nil,
			),
			Text("Hello"),
			nil,
			Text("World"),
		)
		expected := `<div class="container">HelloWorld</div>HelloWorld`
		assert.Render(t, expected, node)
	})
}

func TestEl(t *testing.T) {
	t.Run("El with no children", func(t *testing.T) {
		node := El("div")
		expected := "<div></div>"
		assert.Render(t, expected, node)
	})

	t.Run("El with children", func(t *testing.T) {
		node := El(
			"div",
			Attr("class", "container"),
			Text("Hello"),
			Text("World"),
		)
		expected := `<div class="container">HelloWorld</div>`
		assert.Render(t, expected, node)
	})
}

func TestElVoid(t *testing.T) {
	t.Run("ElVoid with no children", func(t *testing.T) {
		node := ElVoid("img")
		expected := "<img>"
		assert.Render(t, expected, node)
	})

	t.Run("ElVoid with ignored children", func(t *testing.T) {
		node := ElVoid(
			"img",
			Attr("class", "container"),
			Text("Ignored"),
		)
		expected := `<img class="container">`
		assert.Render(t, expected, node)
	})
}

func TestAttr(t *testing.T) {
	t.Run("Attr with value", func(t *testing.T) {
		node := Attr("class", "container")
		expected := `class="container"`
		assert.Render(t, expected, node)
	})

	t.Run("Attr with empty value", func(t *testing.T) {
		node := Attr("class", "")
		expected := `class=""`
		assert.Render(t, expected, node)
	})

	t.Run("Attr with empty name", func(t *testing.T) {
		node := Attr("", "value")
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("Attr with no value at all", func(t *testing.T) {
		node := Attr("disabled")
		expected := `disabled`
		assert.Render(t, expected, node)
	})
}

func TestAttrBool(t *testing.T) {
	t.Run("AttrBool with true value", func(t *testing.T) {
		node := AttrBool("disabled", true)
		expected := `disabled`
		assert.Render(t, expected, node)
	})

	t.Run("AttrBool with false value", func(t *testing.T) {
		node := AttrBool("disabled", false)
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("AttrBool with empty name and true value", func(t *testing.T) {
		node := AttrBool("", true)
		expected := ""
		assert.Render(t, expected, node)
	})

	t.Run("AttrBool inside an element", func(t *testing.T) {
		node := ElVoid("input",
			Attr("type", "checkbox"),
			AttrBool("checked", true),
			AttrBool("disabled", false),
		)
		expected := `<input type="checkbox" checked>`
		assert.Render(t, expected, node)
	})
}

func TestText(t *testing.T) {
	t.Run("Text with content", func(t *testing.T) {
		node := Text("Hello <b>World</b>")
		expected := "Hello &lt;b&gt;World&lt;/b&gt;"
		assert.Render(t, expected, node)
	})

	t.Run("Text with empty content", func(t *testing.T) {
		node := Text("")
		expected := ""
		assert.Render(t, expected, node)
	})
}

func TestTextf(t *testing.T) {
	t.Run("Textf with formatted content", func(t *testing.T) {
		node := Textf("Hello, %s!", "World")
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("Textf with escaped content", func(t *testing.T) {
		node := Textf("Hello <b>%s</b>", "World")
		expected := "Hello &lt;b&gt;World&lt;/b&gt;"
		assert.Render(t, expected, node)
	})

	t.Run("Textf with empty content", func(t *testing.T) {
		node := Textf("")
		expected := ""
		assert.Render(t, expected, node)
	})
}

func TestRaw(t *testing.T) {
	t.Run("Raw with content", func(t *testing.T) {
		node := Raw("<script>alert('XSS');</script>")
		expected := "<script>alert('XSS');</script>"
		assert.Render(t, expected, node)
	})

	t.Run("Raw with empty content", func(t *testing.T) {
		node := Raw("")
		expected := ""
		assert.Render(t, expected, node)
	})
}

func TestRawf(t *testing.T) {
	t.Run("Rawf with formatted content", func(t *testing.T) {
		node := Rawf("<div>%s</div>", "Content")
		expected := "<div>Content</div>"
		assert.Render(t, expected, node)
	})
}

func TestIf(t *testing.T) {
	t.Run("If condition true", func(t *testing.T) {
		node := If(true, Text("Hello"))
		expected := "Hello"
		assert.Render(t, expected, node)
	})

	t.Run("If condition false", func(t *testing.T) {
		node := If(false, Text("Hello"))
		expected := ""
		assert.Render(t, expected, node)
	})
}

func TestIfFunc(t *testing.T) {
	t.Run("IfFunc condition true", func(t *testing.T) {
		node := IfFunc(true, func() Node {
			return Text("Hello")
		})
		expected := "Hello"

		assert.Render(t, expected, node)
	})

	t.Run("IfFunc condition false", func(t *testing.T) {
		node := IfFunc(false, func() Node {
			return Text("Hello")
		})
		expected := ""

		assert.Render(t, expected, node)
	})
}

func TestMap(t *testing.T) {
	t.Run("Map with integers", func(t *testing.T) {
		slice := []int{1, 2, 3}
		node := Map(slice, func(i int) Node {
			return Textf("%d", i)
		})
		expected := "123"

		assert.Render(t, expected, node)
	})

	t.Run("Map with strings", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		node := Map(slice, func(s string) Node {
			return Text(s)
		})
		expected := "abc"

		assert.Render(t, expected, node)
	})

	t.Run("Map with empty slice", func(t *testing.T) {
		slice := []int{}
		node := Map(slice, func(i int) Node {
			return Textf("%d", i)
		})
		expected := ""

		assert.Render(t, expected, node)
	})

	t.Run("Complex map", func(t *testing.T) {
		slice := []int{1, 2, 3, 4}
		node := El("div", Map(slice, func(i int) Node {
			if i == 1 {
				return El("span", Text("Hello"))
			}
			if i == 2 {
				return Attr("class", "test")
			}
			if i == 3 {
				return El("span", Text("World"))
			}
			return El("span", Text("!"))
		}))
		expected := `
			<div class="test">
				<span>Hello</span>
				<span>World</span>
				<span>!</span>
			</div>
		`

		assert.RenderNoSpaces(t, expected, node)
	})

	t.Run("HTML5 Structure", func(t *testing.T) {
		node := Group(
			DocType(),
			Html(
				Head(
					TitleEl(Text("Title")),
				),
				Body(
					H1(Text("Hello, World!")),
				),
			),
		)

		expected := `
			<!DOCTYPE html>
			<html>
				<head>
					<title>Title</title>
				</head>
				<body>
					<h1>Hello, World!</h1>
				</body>
			</html>
		`

		assert.RenderNoSpaces(t, expected, node)
	})

	t.Run("Advanced with components and conditions", func(t *testing.T) {
		button := func(text string, isLarge bool) Node {
			return Button(
				ClassMap{
					"btn":   true,
					"large": isLarge,
				},
				Text(text),
			)
		}

		node := Group(
			Div(
				StyleMap{
					"display: flex": true,
				},

				Main(
					Class("buttons"),
					button("Click me", false),
					button("Click me", true),
				),
			),
		)

		expected := `
			<div style="display: flex">
				<main class="buttons">
					<button class="btn">Click me</button>
					<button class="btn large">Click me</button>
				</main>
			</div>
		`

		assert.RenderNoSpaces(t, expected, node)
	})

	t.Run("Eval", func(t *testing.T) {
		node := Eval(func() Node {
			str := "World"
			return Textf("Hello, %s!", str)
		})
		expected := "Hello, World!"
		assert.Render(t, expected, node)
	})

	t.Run("Eval as child", func(t *testing.T) {
		node := Div(
			Class("container"),
			Eval(func() Node {
				return Text("Hello, World!")
			}),
		)
		expected := `<div class="container">Hello, World!</div>`
		assert.Render(t, expected, node)
	})

	t.Run("Eval returning attribute", func(t *testing.T) {
		node := Div(
			Eval(func() Node {
				return Attr("class", "test")
			}),
			Text("Hello, World!"),
		)
		expected := `<div class="test">Hello, World!</div>`
		assert.Render(t, expected, node)
	})

	t.Run("Nested Eval", func(t *testing.T) {
		node := Div(
			Class("container"),
			Eval(func() Node {
				return Eval(func() Node {
					return Eval(func() Node {
						return Eval(func() Node {
							return Text("Hello, World!")
						})
					})
				})
			}),
		)
		expected := `<div class="container">Hello, World!</div>`
		assert.Render(t, expected, node)
	})
}
