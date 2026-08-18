package features

import (
	"testing"
)

func TestInspirationManager(t *testing.T) {
	im := NewInspirationManager()
	q1 := im.CurrentQuote()
	if q1.Quote == "" || q1.Author == "" {
		t.Fatalf("expected non-empty quote and author, got quote: %q, author: %q", q1.Quote, q1.Author)
	}

	q2 := im.NextQuote()
	if q2.Quote == "" || q2.Author == "" {
		t.Fatalf("expected non-empty quote after NextQuote, got quote: %q, author: %q", q2.Quote, q2.Author)
	}

	if q1.Quote == q2.Quote {
		t.Fatalf("expected NextQuote to cycle quote, but got duplicate: %q", q1.Quote)
	}
}
