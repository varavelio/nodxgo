package nodx

import (
	"fmt"
	"io"
)

// ensure that nodeText implements the Node interface.
var _ Node = (*nodeText)(nil)

// nodeText represents a text node.
//
// IMPORTANT: This node renders the text as is, without escaping it. It is the
// responsibility of the caller to ensure that the text is properly escaped.
type nodeText struct {
	text string
}

// newNodeText creates a new text node.
//
// IMPORTANT: This node renders the text as is, without escaping it. It is the
// responsibility of the caller to ensure that the text is properly escaped.
func newNodeText(text string) nodeText {
	return nodeText{
		text: text,
	}
}

func (nt nodeText) Render(w io.Writer) error {
	if nt.text == "" {
		return nil
	}
	if _, err := io.WriteString(w, nt.text); err != nil {
		return fmt.Errorf("failed to render text: %w", err)
	}

	return nil
}

func (nt nodeText) RenderString() (string, error) {
	return nt.text, nil
}

func (nt nodeText) RenderBytes() ([]byte, error) {
	return []byte(nt.text), nil
}

func (nt nodeText) IsElement() bool {
	return true
}

func (nt nodeText) IsAttribute() bool {
	return false
}

func (nt nodeText) String() string {
	return nt.text
}
