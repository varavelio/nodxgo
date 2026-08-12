package nodx

import (
	"testing"
)

func BenchmarkEscapeHTMLNoSpecialChars(b *testing.B) {
	input := "Hello, world! This is a plain text without special characters."
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EscapeHTML(input)
	}
}

func BenchmarkEscapeHTMLWithSpecialChars(b *testing.B) {
	input := `<a href="https://example.com?a=1&b=2">Tom & Jerry's</a>`
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EscapeHTML(input)
	}
}

func BenchmarkEscapeHTMLUTF8(b *testing.B) {
	input := "café naïve — 日本語 🎉 sin caracteres especiales"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EscapeHTML(input)
	}
}

func BenchmarkRenderElement(b *testing.B) {
	node := Div(
		Class("card", "card-lg"),
		Id("main"),
		Href("https://example.com?a=1&b=2"),
		Text("Hello, <world>!"),
	)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = node.RenderString()
	}
}

func BenchmarkRenderManyChildren(b *testing.B) {
	children := make([]Node, 0, 100)
	for range 100 {
		children = append(children, Li(Class("item"), Text("Item")))
	}
	node := Ul(children...)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = node.RenderString()
	}
}

func BenchmarkRenderClassMap(b *testing.B) {
	node := ClassMap{"a": true, "b": false, "c": true, "d": true}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = node.RenderString()
	}
}

func BenchmarkAttrList(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = AttrList("class", "a", "b", "c")
	}
}
