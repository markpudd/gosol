package gosol

import  (
  "testing"
)

func TestFlipCard(t *testing.T) {
  card := NewCard(20)
  if card.String() != "XXX" {
    t.Errorf("Card should not be initial flipped")
  }
  card.FlipCard()
  if card.String() == "XXX" {
    t.Errorf("Card not flipped")
  }
}
