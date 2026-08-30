package sanitize

import "testing"

func TestPlainStripsANSI(t *testing.T) {
	in := "hello\x1b[31mred\x1b[0m"
	got := Plain(in)
	if got != "hellored" {
		t.Fatalf("Plain(%q) = %q, want hellored", in, got)
	}
}

func TestPlainStripsControlChars(t *testing.T) {
	in := "ok\x07\x08\n\t"
	got := Plain(in)
	if got != "ok\n\t" {
		t.Fatalf("Plain(%q) = %q, want ok\\n\\t", in, got)
	}
}

func TestPlainPreservesNewlines(t *testing.T) {
	in := "line1\nline2"
	if Plain(in) != in {
		t.Fatalf("newlines should be preserved")
	}
}
