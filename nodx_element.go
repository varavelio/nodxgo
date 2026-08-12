package nodx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// ensure that nodeElement implements the Node interface.
var _ Node = (*nodeElement)(nil)

// nodeElement represents an HTML element.
type nodeElement struct {
	isVoid   bool
	name     string
	children []Node
}

// newNodeElement creates a new HTML element.
func newNodeElement(isVoid bool, name string, children ...Node) nodeElement {
	return nodeElement{
		isVoid:   isVoid,
		name:     name,
		children: children,
	}
}

// categorizeChildren categorizes the children into elements and attributes.
//
//   - The nil nodes are ignored.
//   - Groups are expanded so that the nodes in the group become children of the
//     element.
//
// It returns two slices, the first are elements and the second are attributes.
func categorizeChildren(children []Node) ([]Node, []Node) {
	elements := make([]Node, 0, len(children))
	attributes := make([]Node, 0, len(children))

	for _, child := range children {
		if child == nil {
			continue
		}
		if group, ok := child.(nodeGroup); ok {
			els, attrs := categorizeChildren(group)
			elements = append(elements, els...)
			attributes = append(attributes, attrs...)
			continue
		}
		if child.IsElement() {
			elements = append(elements, child)
		}
		if child.IsAttribute() {
			attributes = append(attributes, child)
		}
	}

	return elements, attributes
}

func (ne nodeElement) Render(w io.Writer) error {
	if ne.name == "" {
		return nil
	}

	els, attrs := categorizeChildren(ne.children)

	if _, err := io.WriteString(w, "<"); err != nil {
		return fmt.Errorf("failed to render element name: %w", err)
	}
	if _, err := io.WriteString(w, ne.name); err != nil {
		return fmt.Errorf("failed to render element name: %w", err)
	}

	for _, attr := range attrs {
		// Add a space between each attribute.
		if _, err := io.WriteString(w, " "); err != nil {
			return fmt.Errorf("failed to render space between attributes: %w", err)
		}

		err := attr.Render(w)
		if err != nil {
			return fmt.Errorf("failed to render child attribute: %w", err)
		}
	}

	if _, err := io.WriteString(w, ">"); err != nil {
		return fmt.Errorf("failed to render element end: %w", err)
	}

	if !ne.isVoid {
		for _, el := range els {
			err := el.Render(w)
			if err != nil {
				return fmt.Errorf("failed to render child element: %w", err)
			}
		}

		if _, err := io.WriteString(w, "</"); err != nil {
			return fmt.Errorf("failed to render element end: %w", err)
		}
		if _, err := io.WriteString(w, ne.name); err != nil {
			return fmt.Errorf("failed to render element end: %w", err)
		}
		if _, err := io.WriteString(w, ">"); err != nil {
			return fmt.Errorf("failed to render element end: %w", err)
		}
	}

	return nil
}

func (ne nodeElement) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := ne.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (ne nodeElement) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := ne.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (ne nodeElement) IsElement() bool {
	return true
}

func (ne nodeElement) IsAttribute() bool {
	return false
}

func (ne nodeElement) String() string {
	str, _ := ne.RenderString()
	return str
}
