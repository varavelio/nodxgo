package nodx

import (
	"strings"
	"testing"
)

// FuzzEscapeHTML ensures that escaping never leaves raw special characters
// behind, is reversible and preserves non-ASCII (UTF-8) bytes exactly.
func FuzzEscapeHTML(f *testing.F) {
	f.Add("plain text")
	f.Add(`<script>alert("xss")</script>`)
	f.Add(`Tom & Jerry's <b>"bold"</b>`)
	f.Add("café naïve — 日本語 🎉")
	f.Add(`&amp;&lt;&gt;&quot;&#39;`)
	f.Add("")
	f.Add(strings.Repeat("&", 100))

	f.Fuzz(func(t *testing.T, input string) {
		got := EscapeHTML(input)

		// No raw special characters may remain in the output.
		if strings.ContainsAny(got, "<>\"'") {
			t.Fatalf("raw special character left unescaped: %q", got)
		}

		// Every ampersand in the output must introduce a known entity.
		for i := 0; i < len(got); i++ {
			if got[i] != '&' {
				continue
			}
			rest := got[i:]
			if !strings.HasPrefix(rest, "&amp;") &&
				!strings.HasPrefix(rest, "&lt;") &&
				!strings.HasPrefix(rest, "&gt;") &&
				!strings.HasPrefix(rest, "&quot;") &&
				!strings.HasPrefix(rest, "&#39;") {
				t.Fatalf("unknown entity in output: %q", got[i:])
			}
		}

		// Escaping must be reversible.
		if unescapeHTML(got) != input {
			t.Fatalf(
				"round-trip mismatch: input %q -> output %q -> %q",
				input,
				got,
				unescapeHTML(got),
			)
		}

		// Non-ASCII bytes must be preserved exactly.
		if nonASCII(input) != nonASCII(got) {
			t.Fatalf("non-ASCII bytes changed: input %q -> output %q", input, got)
		}
	})
}

// FuzzRender ensures that rendering arbitrary nodes never panics and is
// deterministic.
func FuzzRender(f *testing.F) {
	f.Add("div", "class", "text")
	f.Add("script", "onclick", `<img src=x onerror=alert(1)>`)
	f.Add("", "", "")
	f.Add("日本語", "attr", "value")

	f.Fuzz(func(t *testing.T, el, attr, text string) {
		node := El(el, Attr(attr, text), Text(text), Raw(text))

		first, err1 := node.RenderString()
		second, err2 := node.RenderString()
		if err1 != nil || err2 != nil {
			t.Fatalf("unexpected error: %v / %v", err1, err2)
		}
		if first != second {
			t.Fatalf("non-deterministic render: %q vs %q", first, second)
		}

		if el != "" && !strings.Contains(first, el) {
			t.Fatalf("element name missing from output: %q", first)
		}
	})
}

// unescapeHTML reverses the escaping done by EscapeHTML.
func unescapeHTML(s string) string {
	return strings.NewReplacer(
		"&amp;", "&",
		"&lt;", "<",
		"&gt;", ">",
		"&quot;", "\"",
		"&#39;", "'",
	).Replace(s)
}

// nonASCII returns the subsequence of non-ASCII bytes of s.
func nonASCII(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}
