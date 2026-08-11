// Package nodx is a modern and developer-friendly Go template engine for
// generating safe, clean, and maintainable HTML. Designed for maximum
// productivity and easy maintenance, it combines simplicity, type safety
// and robust formatting, making it the perfect fit for your Go-based web
// projects.
//
// In NodX, everything is a node and anything that implements the Node interface
// can be rendered as HTML or used as a Node child.
package nodx

import (
	"fmt"
	"io"
)

// Node is the interface that wraps the basic Render methods used in the
// different NodX node types.
//
// Anything that implements this interface can be used as a node.
type Node interface {
	// Render writes the node HTML to the writer.
	Render(w io.Writer) error

	// RenderString returns the node HTML as a string.
	RenderString() (string, error)

	// RenderBytes returns the node HTML as a byte slice.
	RenderBytes() ([]byte, error)

	// IsElement indicates if the node should be rendered as an element.
	IsElement() bool

	// IsAttribute indicates if the node should be rendered as an attribute.
	IsAttribute() bool

	// String returns the node HTML as a string and implements the fmt.Stringer interface.
	//
	// This is only for convenience and debugging purposes and it ignores the
	// error return value.
	//
	// You should use RenderString() whenever possible.
	String() string
}

// Group combines multiple nodes into a single node without wrapping them in any HTML tag.
//
// When rendered directly, it will call Render on all the nodes in the group
// sequentially.
//
// When used as a child of another node, it will be expanded so that the nodes
// in the group become children of the group's parent.
func Group(nodes ...Node) Node {
	return newNodeGroup(nodes...)
}

// El creates a new Node representing an HTML element with the given name.
func El(name string, children ...Node) Node {
	return newNodeElement(false, name, children...)
}

// ElVoid creates a new Node representing an HTML void element with the given name.
// Void elements are self-closing, such as <img> or <input>.
func ElVoid(name string, children ...Node) Node {
	return newNodeElement(true, name, children...)
}

// Attr creates a new Node representing an HTML attribute with a name and value.
// The value is HTML-escaped to prevent XSS attacks.
//
// If you don't provide a value, the attribute will be rendered without it.
func Attr(name string, value ...string) Node {
	return newNodeAttribute(name, value...)
}

// Text creates a new Node representing an escaped HTML text node.
// The value is HTML-escaped to prevent XSS attacks.
func Text(value string) Node {
	return newNodeText(EscapeHTML(value))
}

// Textf creates a new Node representing an escaped HTML text node.
// The value is formatted with fmt.Sprintf and then HTML-escaped to prevent XSS attacks.
func Textf(format string, a ...any) Node {
	return Text(fmt.Sprintf(format, a...))
}

// Raw creates a new Node representing a raw HTML text node.
// The value is not escaped, so the caller must ensure the content is safe.
// Useful for rendering raw HTML, like <script> or <style> tags.
func Raw(value string) Node {
	return newNodeText(value)
}

// Rawf creates a new Node representing a raw HTML text node.
// The value is formatted with fmt.Sprintf and not escaped, so ensure its safety.
// Useful for rendering raw HTML, like <script> or <style> tags.
func Rawf(format string, a ...any) Node {
	return Raw(fmt.Sprintf(format, a...))
}

// If renders a Node based on the provided boolean condition.
// If the condition is true, the node is rendered; otherwise, it renders nothing.
func If(condition bool, node Node) Node {
	if condition {
		return node
	}
	return Group()
}

// IfFunc executes and renders the result of a function based on the provided boolean condition.
// If the condition is true, the function is executed and rendered; otherwise, nothing is executed nor rendered.
func IfFunc(condition bool, function func() Node) Node {
	if condition {
		return function()
	}
	return Group()
}

// Map transforms a slice of any type into a group of nodes by applying a function to each element.
// Returns a single Node that contains all the resulting nodes.
func Map[T any](slice []T, function func(T) Node) Node {
	nodes := make([]Node, len(slice))
	for i, item := range slice {
		nodes[i] = function(item)
	}
	return Group(nodes...)
}

// Eval executes a provided function and integrates its resulting Node into the current node tree.
//
// Use Eval to insert dynamic content, apply complex logic, or generate nodes on the fly.
//
// Example:
//
//	node := nodx.Group(
//		nodx.Div(
//		  nodx.Class("container"),
//		  nodx.Eval(func() nodx.Node {
//		    if condition {
//		      return nodx.Text("Condition is true")
//		    }
//		    return nodx.Text("Condition is false")
//		  }),
//		),
//	)
func Eval(fn func() Node) Node {
	return fn()
}

// DocType Defines the document type.
//
// This is a special element and should be used inside a nodx.Group.
//
// Output: <!DOCTYPE html>.
func DocType() Node {
	return Raw("<!DOCTYPE html>")
}
