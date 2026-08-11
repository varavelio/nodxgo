package nodx

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// ensure that StyleMap implements the Node interface.
var _ Node = (*StyleMap)(nil)

// StyleMap represents a map of CSS style rules with conditional rendering.
// The keys are the complete CSS rules (e.g., "border: 1px solid black"), and
// the values are boolean conditions indicating whether each rule should be
// included in the final output.
//
// Example:
//
//	sm := StyleMap{
//		"border: 1px solid black": true,  // Included
//		"padding: 10px":           false, // Excluded
//		"margin: 5px":             true,  // Included
//	}
//
// This will render the style attribute as: style="border: 1px solid black; margin: 5px".
type StyleMap map[string]bool

func (sm StyleMap) Render(w io.Writer) error {
	styles := []string{}

	for style, condition := range sm {
		if condition {
			styles = append(styles, style)
		}
	}

	sort.Strings(styles)
	stylesStr := strings.Join(styles, "; ")

	err := Attr("style", stylesStr).Render(w)
	if err != nil {
		return fmt.Errorf("failed to render StyleMap: %w", err)
	}

	return nil
}

func (sm StyleMap) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := sm.Render(buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (sm StyleMap) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := sm.Render(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (sm StyleMap) IsElement() bool {
	return false
}

func (sm StyleMap) IsAttribute() bool {
	return true
}

func (sm StyleMap) String() string {
	str, _ := sm.RenderString()
	return str
}
