package nodx

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// ensure that ClassMap implements the Node interface.
var _ Node = (*ClassMap)(nil)

// ClassMap represents a map of class names with conditional rendering.
// The keys are the class names, and the values are boolean expressions
// (conditions) that determine whether each class should be included
// in the final output.
//
// Example:
//
//	isOdd := func(n int) bool { return n % 2 != 0 }
//	isEven := func(n int) bool { return n % 2 == 0 }
//
//	renderOdd := isOdd(3)   // true
//	renderEven := isEven(3) // false
//
//	cm := ClassMap{
//		"odd-class":  renderOdd,      // Included
//		"even-class": renderEven,     // Excluded
//		"always-on":  true,           // Always included
//	}
//
// This will render the class attribute as: class="odd-class always-on".
type ClassMap map[string]bool

func (cm ClassMap) Render(w io.Writer) error {
	classes := make([]string, 0, len(cm))

	for class, condition := range cm {
		if condition {
			classes = append(classes, class)
		}
	}

	sort.Strings(classes)
	classesStr := strings.Join(classes, " ")

	err := Attr("class", classesStr).Render(w)
	if err != nil {
		return fmt.Errorf("failed to render ClassMap: %w", err)
	}

	return nil
}

func (cm ClassMap) RenderString() (string, error) {
	buf := &strings.Builder{}
	err := cm.Render(buf)
	return buf.String(), err
}

func (cm ClassMap) RenderBytes() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := cm.Render(buf)
	return buf.Bytes(), err
}

func (cm ClassMap) IsElement() bool {
	return false
}

func (cm ClassMap) IsAttribute() bool {
	return true
}

func (cm ClassMap) String() string {
	str, _ := cm.RenderString()
	return str
}
