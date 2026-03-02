package engine

import (
	"testing"
	"github.com/deng37/grab-your-labubu/model"
)

func TestGrabItem(t *testing.T) {
	// Setup: Stock 1 only
	store := &model.LabubuStore{
		StockName: "Labubu Test",
		Count:     1,
	}

	// Test 1: Success
	if success := GrabItem(store); !success {
		t.Errorf("Grab first stock - should success")
	}

	// Test 2: Failed, out of stock
	if success := GrabItem(store); success {
		t.Errorf("Grab another - should failed because out of stock")
	}

	// Test 3: Check stock
	if store.Count < 0 {
		t.Errorf("Error: Minus stock! Remaining: %d", store.Count)
	}
}

func BenchmarkGrabItem(b *testing.B) {
	store := &model.LabubuStore{StockName: "Bench", Count: b.N}

	for i := 0; i < b.N; i++ {
		GrabItem(store)
	}
}