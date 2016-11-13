package gosol

import  (
  "testing"
)

//  This is a super basic test.   As the random seed is always
// the same the first card should differ
func TestShuffle(t *testing.T) {
  deck := NewDeck(52)
  card := deck.Cards[0]
  deck.Shuffle()
  if card == deck.Cards[0] {
    t.Errorf("Deck not shuffled")
  }
}

func TestTopCard(t *testing.T) {
  deck := NewDeck(52)
  card := deck.TopCard()
  if card != deck.Cards[51] {
    t.Errorf("Top card does not equal top of deck")
  }
  if len(deck.Cards) != 52 {
    t.Errorf("Deck length is %d should be 52", len(deck.Cards))
  }
}

func TestTakeCard(t *testing.T) {
  deck := NewDeck(52)
  card := deck.TakeCard()
  if card ==  deck.TopCard() {
    t.Errorf("Top card equals top of deck")
  }
  if card.card != 51 {
    t.Errorf("Card is %d but should be 51",card.card)
  }
  if len(deck.Cards) != 51 {
    t.Errorf("Deck length is %d should be 51", len(deck.Cards))
  }
}


func TestAddCard(t *testing.T) {
  deck := NewDeck(0)
  card := NewCard(0)
  deck.AddCard(card)
  if card !=  deck.TopCard() {
    t.Errorf("Top card is not created card")
  }

}

func TestFlipDeck(t *testing.T) {
  deck1 := NewDeck(0)
  deck2 := NewDeck(0)
  aceClubs := NewCard(0)
  aceClubs.FlipCard()
  deck1.AddCard(aceClubs)
  aceHearts := NewCard(1)
  aceHearts.FlipCard()
  deck1.AddCard(aceHearts)
  twoHearts := NewCard(5)
  twoHearts.FlipCard()
  deck1.AddCard(twoHearts)

  deck1.FlipDeckTo(deck2)
  if len(deck1.Cards) != 0 {
    t.Errorf("Deck length is %d should be 0", len(deck1.Cards))
  }
  if len(deck2.Cards) != 3 {
    t.Errorf("Deck length is %d should be 3", len(deck2.Cards))
  } else {
    if deck2.Cards[0] != twoHearts  {
      t.Errorf("Card 0 is not two Hearts")
    }
    if deck2.Cards[1] != aceHearts  {
      t.Errorf("Card 1 is not ace hearts")
    }
    if deck2.Cards[2] != aceClubs  {
      t.Errorf("Card 2 is not ace Clubs")
    }
    for i:=0;i<3;i++ {
      if deck2.Cards[i].turned {
        t.Errorf("Card %d is turned",i)
      }
    }

  }

}
