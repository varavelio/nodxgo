package nodx

import (
	"testing"

	"github.com/varavelio/nodxgo/internal/assert"
)

func TestEscapeHTML(t *testing.T) {
	t.Run("Escape HTML", func(t *testing.T) {
		input := "<script>alert('XSS')</script>"
		expected := "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;"
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})

	t.Run("All characters", func(t *testing.T) {
		input := "test-&<>'\"-test"
		expected := "test-&amp;&lt;&gt;&#39;&quot;-test"
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})

	t.Run("No escape", func(t *testing.T) {
		input := "test"
		expected := "test"
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})

	t.Run("Empty string", func(t *testing.T) {
		input := ""
		expected := ""
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})

	t.Run("Escape with special characters", func(t *testing.T) {
		input := "test👌áéíóú-&<>'\"-👌test"
		expected := "test👌áéíóú-&amp;&lt;&gt;&#39;&quot;-👌test"
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})

	t.Run("No escape with special characters", func(t *testing.T) {
		input := "test👌✪áéíóúÖ-привет-नमस्ते-你好-مرحبًا"
		expected := "test👌✪áéíóúÖ-привет-नमस्ते-你好-مرحبًا"
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})

	t.Run("Single special character", func(t *testing.T) {
		input := "👌"
		expected := "👌"
		got := EscapeHTML(input)
		assert.Equal(t, expected, got)
	})
}
