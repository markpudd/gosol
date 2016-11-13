package gosol


type Pile []*Card

func CreatePile() Pile {
  return make([]*Card, 0, 52)
}

func (p *Pile) AddCard(card *Card) bool {
  *p = append(*p,card)
  return true
}

func (p *Pile) AddPlayCard(card *Card) bool {
  sl := len(*p)
  if sl >0 && (*p)[sl-1].card&0x1 == card.card&0x1 {
    return false
  }
  if sl >0 && (*p)[sl-1].card>>2 != (card.card>>2)+1 {
    return false
  }
  *p = append(*p,card)
  return true
}


func (p *Pile) AddEndCard(card *Card) bool {
  sl := len(*p)
  if sl >0 {
    if (*p)[sl-1].card&0x3 != card.card&0x3 {
      return false
    }
    if (*p)[sl-1].card>>2 != (card.card>>2)-1 {
      return false
    }
  } else if card.card>>2 != 0 {
    return false
  }
  *p = append(*p,card)
  return true
}

func (p *Pile) RemoveLastCard() *Card {
  if len(*p) == 0 {
    return nil
  }
  ret := (*p)[len(*p)-1]
  *p = (*p)[:len(*p)-1]
  if len(*p) != 0 {
    (*p)[len(*p)-1].turned=true
  }
  return ret
}

func (p *Pile) LastCard() *Card {
  if len(*p) == 0 {
    return nil
  }
  return (*p)[len(*p)-1]
}

func (p *Pile) TurnTop() {
  if len(*p) != 0 {
    (*p)[len(*p)-1].FlipCard()
  }
}

func (p *Pile) MoveStack(pile2 *Pile) bool {
  i := 0
  for i < len(*p) {
    if (*p)[i].turned {
      rm := i
      card := (*p)[i]
      if !pile2.AddPlayCard(card) {
        return false
      }
      y :=i+1
      for y < len(*p) {
        card = (*p)[y]
        pile2.AddCard(card)
        y=y+1
      }
      *p = (*p)[:rm]
      if len(*p) != 0 {
        (*p)[len(*p)-1].turned=true
      }
      return true
    }
    i=i+1
  }
  return false
}
