package features

import (
	"testing"
)

func TestBeverageManager(t *testing.T) {
	bm := NewBeverageManager()

	if bm == nil {
		t.Fatal("expected non-nil BeverageManager")
	}

	cur := bm.CurrentDrink()
	if cur.Type != DrinkDripCoffee {
		t.Errorf("expected default drink to be DrinkDripCoffee, got %v", cur.Type)
	}

	if bm.Warmth != 100 {
		t.Errorf("expected initial warmth to be 100, got %d", bm.Warmth)
	}

	// Test NextDrink
	next := bm.NextDrink()
	if next.Type != DrinkMatchaLatte {
		t.Errorf("expected next drink to be DrinkMatchaLatte, got %v", next.Type)
	}

	// Test TakeSip
	msg, warmth := bm.TakeSip()
	if msg == "" {
		t.Errorf("expected non-empty sip feedback message")
	}
	if warmth != 97 {
		t.Errorf("expected warmth after 1 sip to be 97, got %d", warmth)
	}
	if bm.SipCount != 1 {
		t.Errorf("expected SipCount to be 1, got %d", bm.SipCount)
	}

	// Test BrewDrink
	brewed := bm.BrewDrink(int(DrinkEarlGrey))
	if brewed.Type != DrinkEarlGrey {
		t.Errorf("expected brewed drink to be DrinkEarlGrey, got %v", brewed.Type)
	}
	if bm.Warmth != 100 {
		t.Errorf("expected warmth to reset to 100 on brew, got %d", bm.Warmth)
	}

	// Test PrevDrink
	prev := bm.PrevDrink()
	if prev.Type != DrinkMatchaLatte {
		t.Errorf("expected prev drink to be DrinkMatchaLatte, got %v", prev.Type)
	}

	// Test StatusBadge
	badge := bm.StatusBadge()
	if badge == "" {
		t.Errorf("expected non-empty status badge")
	}
}
