package nodx

import (
	"errors"
	"testing"
)

// failWriter fails on every write.
type failWriter struct{}

func (failWriter) Write(p []byte) (int, error) {
	return 0, errors.New("write failed")
}

func TestRenderErrors(t *testing.T) {
	tests := []struct {
		name string
		node Node
	}{
		{"element", El("div", Id("main"), Text("hello"))},
		{"element with children", El("div", El("span", Text("x")))},
		{"void element", ElVoid("img", Href("x"))},
		{"attribute with value", Attr("class", "a")},
		{"attribute without value", Attr("disabled")},
		{"text", Text("hello")},
		{"group", Group(Text("a"), Text("b"))},
		{"class map", ClassMap{"a": true}},
		{"style map", StyleMap{"color: red": true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.node.Render(failWriter{}); err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
