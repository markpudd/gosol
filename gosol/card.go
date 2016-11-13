package gosol

import (
  "fmt"
)

type Card struct {
  card int
  turned bool
}

func CardLongString(card int) string {
  if card < 0 || card > 51 {
      return "X"
  } else {
    var colour =  [...]string{"black","red"}
    var suit = [...]string{"clubs","hearts","spades","diamonds"}

    return fmt.Sprint((card>>2)+2,suit[card&0x3],colour[card&0x1])
   }
}

func CardShortString(card int) string {
  if card < 0 || card > 51  {
      return "X"
  } else {
    var suit = [...]string{"C","H","S","D"}
    var numbers = [...]string{"A","2","3","4","5","6","7","8","9","T","J","Q","K"}
    if card&0x1 == 1 {
      return fmt.Sprintf("\x1B[31m%s%sR\033[0m",numbers[(card>>2)],suit[card&0x3])
    } else {
      return fmt.Sprintf("%s%sB",numbers[(card>>2)],suit[card&0x3])
    }
   }
}


func NewCard(card int) *Card {
  c := new(Card)
  c.card = card
  c.turned = false
  return c
}

func (gc *Card) FlipCard() {
  gc.turned = !gc.turned
}


func (gc *Card) String() string {
  if gc.turned {
    return CardShortString(gc.card)
  } else {
    return "XXX"
  }
}
