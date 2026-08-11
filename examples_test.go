package nodx_test

import (
	"fmt"

	nodx "github.com/varavelio/nodxgo"
)

func ExampleGroup() {
	node := nodx.Group(
		nodx.DocType(),
		nodx.Html(
			nodx.Head(
				nodx.TitleEl(nodx.Text("Hello, World!")),
			),
			nodx.Body(nodx.Text("Hello, World!")),
		),
	)

	fmt.Println(node)
	// Output:
	// <!DOCTYPE html><html><head><title>Hello, World!</title></head><body>Hello, World!</body></html>
}

func ExampleEl() {
	node := nodx.El("my-custom-element",
		nodx.Text("Hello, World!"),
	)

	fmt.Println(node)
	// Output:
	// <my-custom-element>Hello, World!</my-custom-element>
}

func ExampleElVoid() {
	node := nodx.ElVoid("my-custom-element",
		nodx.Class("container"),
	)

	fmt.Println(node)
	// Output:
	// <my-custom-element class="container">
}

func ExampleAttr() {
	node := nodx.Div(
		nodx.Attr("my-custom-attr", "foo"),
	)

	fmt.Println(node)
	// Output:
	// <div my-custom-attr="foo"></div>
}

func ExampleText() {
	node := nodx.Div(
		nodx.Text("<script>alert('Hello, World!')</script>"),
	)

	fmt.Println(node)
	// Output:
	// <div>&lt;script&gt;alert(&#39;Hello, World!&#39;)&lt;/script&gt;</div>
}

func ExampleTextf() {
	node := nodx.Div(
		nodx.Textf("<script>alert('Hello, %s!')</script>", "World"),
	)

	fmt.Println(node)
	// Output:
	// <div>&lt;script&gt;alert(&#39;Hello, World!&#39;)&lt;/script&gt;</div>
}

func ExampleRaw() {
	node := nodx.Div(
		nodx.Raw("<script>alert('Hello, World!')</script>"),
	)

	fmt.Println(node)
	// Output:
	// <div><script>alert('Hello, World!')</script></div>
}

func ExampleRawf() {
	node := nodx.Div(
		nodx.Rawf("<script>alert('Hello, %s!')</script>", "World"),
	)

	fmt.Println(node)
	// Output:
	// <div><script>alert('Hello, World!')</script></div>
}

func ExampleIf() {
	condition := 3 > 2

	node := nodx.Div(
		nodx.Class("container"),
		nodx.If(condition, nodx.Text("Condition is true")),
	)

	fmt.Println(node)
	// Output:
	// <div class="container">Condition is true</div>
}

func ExampleIfFunc() {
	condition := 3 > 2

	node := nodx.Div(
		nodx.Class("container"),
		nodx.IfFunc(condition, func() nodx.Node {
			// You can add your own go code here
			return nodx.Text("Condition is true")
		}),
	)

	fmt.Println(node)
	// Output:
	// <div class="container">Condition is true</div>
}

func ExampleMap() {
	items := []string{"Item 1", "Item 2", "Item 3"}

	node := nodx.Ul(
		nodx.Map(items, func(item string) nodx.Node {
			return nodx.Li(nodx.Text(item))
		}),
	)

	fmt.Println(node)
	// Output:
	// <ul><li>Item 1</li><li>Item 2</li><li>Item 3</li></ul>
}

func ExampleEval() {
	condition := 3 > 2

	node := nodx.Div(
		nodx.Class("container"),
		nodx.Eval(func() nodx.Node {
			// You can add your own go code here
			if condition {
				return nodx.Text("Condition is true")
			}
			return nodx.Text("Condition is false")
		}),
	)

	fmt.Println(node)
	// Output:
	// <div class="container">Condition is true</div>
}

func ExampleClassMap() {
	shouldHide := 3 > 2

	node := nodx.Div(
		nodx.ClassMap{
			"container": true,
			"other":     false,
			"hidden":    shouldHide,
		},
		nodx.Text("Hello, World!"),
	)

	fmt.Println(node)
	// Output:
	// <div class="container hidden">Hello, World!</div>
}

func ExampleStyleMap() {
	shouldHide := 3 > 2

	node := nodx.Div(
		nodx.StyleMap{
			"color: red":             true,
			"border: 1px solid blue": false,
			"display: none":          shouldHide,
		},
		nodx.Text("Hello, World!"),
	)

	fmt.Println(node)
	// Output:
	// <div style="color: red; display: none">Hello, World!</div>
}
