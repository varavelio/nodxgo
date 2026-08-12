package nodx

import (
	"errors"
	"fmt"
	"testing"
)

// countingWriter counts the writes it receives.
type countingWriter struct{ writes int }

func (w *countingWriter) Write(p []byte) (int, error) {
	w.writes++
	return len(p), nil
}

// failAfterWriter succeeds for n writes, then fails on the next one.
type failAfterWriter struct {
	n      int
	writes int
}

func (w *failAfterWriter) Write(p []byte) (int, error) {
	if w.writes >= w.n {
		return 0, errors.New("write failed")
	}
	w.writes++
	return len(p), nil
}

func TestRenderErrors(t *testing.T) {
	nodes := []Node{
		El("div", Id("main"), Text("hello")),
		El("div", El("span", Text("x"))),
		ElVoid("img", Href("x")),
		Attr("class", "a"),
		Attr("disabled"),
		Text("hello"),
		Group(Text("a"), Text("b")),
		ClassMap{"a": true},
		StyleMap{"color: red": true},
	}

	for _, node := range nodes {
		t.Run(fmt.Sprintf("%T", node), func(t *testing.T) {
			// The writer never fails: count how many writes the node performs.
			cw := &countingWriter{}
			if err := node.Render(cw); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Failing at any single write must propagate an error.
			for i := range cw.writes {
				fw := &failAfterWriter{n: i}
				if err := node.Render(fw); err == nil {
					t.Errorf("expected error when write %d fails", i)
				}
			}
		})
	}
}
