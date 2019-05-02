package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length 16, but got %v", len(d))
	}

	if d[0] != "Ace of Diamonds" {
		t.Errorf("Ecpected first deck Ace of Diamonds, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Ecpected last deck Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveAndRead(t *testing.T) {
	os.Remove("_decktesting")
	d := createDeck()
	d.save("_decktesting")
	newDeck := read("_decktesting")

	if len(newDeck) != len(d) {
		t.Errorf("Ecpected got length 16, but got %v", len(newDeck))
	}
}
