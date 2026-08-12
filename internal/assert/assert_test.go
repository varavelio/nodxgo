package assert

import (
	"errors"
	"io"
	"testing"
)

// TestAssertionsPass exercises the success paths of all assertion helpers.
//
// The failure paths (the t.Errorf/t.Error calls) are intentionally not covered
// here: executing them would fail the test by design. They are the reporting
// side of the helpers and are exercised whenever a real assertion fails.
func TestAssertionsPass(t *testing.T) {
	NoError(t, nil)
	Error(t, errors.New("boom"))
	Equal(t, 1, 1)
	Equal(t, "a", "a")
	NotEqual(t, 1, 2)
	True(t, true)
	False(t, false)
	Nil(t, nil)
	NotNil(t, "x")
}

// testNode is a minimal Node implementation for Render/RenderNoSpaces.
type testNode struct{}

func (testNode) Render(w io.Writer) error {
	_, err := w.Write([]byte("<div>Hello</div>"))
	return err
}

func (testNode) RenderString() (string, error) {
	return "<div>Hello</div>", nil
}

func (testNode) RenderBytes() ([]byte, error) {
	return []byte("<div>Hello</div>"), nil
}

func (testNode) String() string {
	return "<div>Hello</div>"
}

func TestRender(t *testing.T) {
	Render(t, "<div>Hello</div>", testNode{})
}

func TestRenderNoSpaces(t *testing.T) {
	RenderNoSpaces(t, "<div> Hello </div>", testNode{})
}
