package blog

import (
	"strings"
	"testing"
)

func TestFormatMarkdownHeadings(t *testing.T) {
	out := FormatMarkdown("# Hello\n\nBody text.", 80)
	if !strings.Contains(out, "Hello") {
		t.Fatalf("expected heading rendered: %q", out)
	}
	if !strings.Contains(out, "Body text") {
		t.Fatalf("expected body rendered: %q", out)
	}
}

func TestFormatMarkdownSkipsImages(t *testing.T) {
	out := FormatMarkdown("![alt](http://example.com/x.png)\n\nAfter.", 80)
	if strings.Contains(out, "![") {
		t.Fatalf("expected image markdown stripped: %q", out)
	}
	if !strings.Contains(out, "After") {
		t.Fatalf("expected content after image: %q", out)
	}
}
