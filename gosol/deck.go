package gosol

import (
  "math/rand"
)



type Deck struct {
  Cards []*Card
}

func NewDeck(size int) *Deck {
  d := new(Deck)
  d.Cards =  make([]*Card, size, 52) //new(Deck)
  for i := 0; i < size; i++ {
    d.Cards[i]=NewCard(i)
  }
  return d;
}


func (d *Deck) Shuffle() {
  for i := 0; i < 52; i++ {
    r := rand.Intn(52)
    tmp := d.Cards[i]
    d.Cards[i] = d.Cards[r]
    d.Cards[r] = tmp
  }
}

func (d *Deck) TopCard() *Card {
  if len(d.Cards) == 0 {
    return nil
  } else {
    return d.Cards[len(d.Cards)-1]
  }
}

func (d *Deck) TakeCard() *Card {
    if len(d.Cards) == 0 {
      return nil
    } else {
      c := d.Cards[len(d.Cards)-1]
      d.Cards = d.Cards[:len(d.Cards)-1]
      return c
    }
}

func (d *Deck) AddCard(card *Card) {
  d.Cards = append(d.Cards,card)
}

func (d *Deck) FlipDeckTo(deck2 *Deck) {
    for i := len(d.Cards)-1;i>=0;i-- {
      card  := d.Cards[i]
      card.turned = false
      deck2.AddCard(card)
    }
    d.Cards = d.Cards[:0]
}
