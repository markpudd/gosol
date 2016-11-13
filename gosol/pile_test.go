package gosol

import  (
  "testing"
)

func TestAddPileCard(t *testing.T) {
  pile := CreatePile()

  aceClubs := NewCard(0)
  aceSpades := NewCard(2)
  if !pile.AddCard(aceClubs) {
      t.Errorf("Unbale to add start card")
  }
  if !pile.AddCard(aceSpades) {
      t.Errorf("Unbale to add second card")
  }
}

func TestAddPlayCard(t *testing.T) {
  pile := CreatePile()

  aceClubs := NewCard(0)
  aceHearts := NewCard(1)
  twoHearts := NewCard(5)
  if !pile.AddPlayCard(twoHearts) {
      t.Errorf("Unable to add start card")
  }
  if pile.AddPlayCard(aceHearts) {
      t.Errorf("Able to add illegal card")
  }
  if !pile.AddPlayCard(aceClubs) {
      t.Errorf("Unable to add legal card")
  }
  if pile.AddPlayCard(twoHearts) {
      t.Errorf("Able to add illegal card")
  }
  if len(pile) != 2 {
    t.Errorf("Pile length %d but should be 2",len(pile))
  }
}


func TestAddPlayCardEnd(t *testing.T) {
  pile := CreatePile()

  aceClubs := NewCard(0)
  twoHearts := NewCard(5)
  twoClubs := NewCard(4)

  if pile.AddEndCard(twoClubs) {
      t.Errorf("Able to add inital end that is not ace %d", twoClubs.card>>2 )
  }

  if !pile.AddEndCard(aceClubs) {
      t.Errorf("Unable to add inital end card")
  }

  if pile.AddEndCard(twoHearts) {
      t.Errorf("Able to add 2 hearts to ace of clubs")
  }
  if !pile.AddEndCard(twoClubs) {
      t.Errorf("Unable to add two clubs to ace of clubs")
  }
  if len(pile) != 2 {
    t.Errorf("Pile length %d but should be 2",len(pile))
  }
}

func TestRemoveLastCard(t *testing.T) {
  pile := CreatePile()
  aceClubs := NewCard(0)
  twoHearts := NewCard(5)
  pile.AddPlayCard(twoHearts)
  pile.AddPlayCard(aceClubs)
  if len(pile) != 2 {
    t.Errorf("Pile length %d but should be 2",len(pile))
  }
  if pile.RemoveLastCard() != aceClubs  {
    t.Errorf("Did not get ace of clubs")
  }
  if len(pile) != 1 {
    t.Errorf("Pile length %d but should be 1",len(pile))
  }
  if pile.RemoveLastCard() != twoHearts {
        t.Errorf("Did not get two of Hearts")

  }
  if pile.RemoveLastCard() != nil {
    t.Errorf("Did not get nil")
  }
}

func TestLastCard(t *testing.T) {
  pile := CreatePile()
  aceClubs := NewCard(0)
  twoHearts := NewCard(5)
  pile.AddCard(aceClubs)
  pile.AddCard(twoHearts)
  if len(pile) != 2 {
    t.Errorf("Pile length %d but should be 2",len(pile))
  }
  if pile.LastCard() != twoHearts {
    t.Errorf("Did not get two of Hearts")
  }
  if len(pile) != 2 {
    t.Errorf("Pile length %d but should be 2",len(pile))
  }
}

func TestMoveStack(t *testing.T) {
  pile1 := CreatePile()
  pile2 := CreatePile()

  randCard := NewCard(43)
  twoClubs := NewCard(4)
  aceHearts := NewCard(3)
  aceHearts.FlipCard()
  twoClubs.FlipCard()

  threeHearts := NewCard(9)


  pile1.AddCard(randCard)
  pile1.AddCard(twoClubs)
  pile1.AddCard(aceHearts)

  pile2.AddCard(threeHearts)

  if !pile1.MoveStack(&pile2) {
    t.Errorf("Unable to move pile")
  }
  if len(pile1) !=1 {
    t.Errorf("Pile has %d but should have 1 card",len(pile1))
  }
  if len(pile2) !=3 {
    t.Errorf("Pile has %d but should have 3 card",len(pile2))
  }

  cc:=pile2.RemoveLastCard()
  if cc != aceHearts {
    t.Errorf("Card not aceHearts but is %s",cc.String())
  }
  cc=pile2.RemoveLastCard()
  if cc  != twoClubs {
    t.Errorf("Card not twoClubs but is %s",cc.String())
  }
  cc=pile2.RemoveLastCard()
  if cc  != threeHearts   {
  t.Errorf("Card not threeHearts but is %s",cc.String())
  }

  cc=pile1.RemoveLastCard()
  if cc  != randCard   {
  t.Errorf("Card not %s but is %s",randCard.String(),cc.String())
  }
}

func TestMoveStack2(t *testing.T) {
  pile1 := CreatePile()
  pile2 := CreatePile()

  twoClubs := NewCard(4)
  twoClubs.FlipCard()

  aceHearts := NewCard(3)
  aceHearts.FlipCard()

  pile1.AddCard(aceHearts)
  pile2.AddCard(twoClubs)

  if !pile1.MoveStack(&pile2) {
    t.Errorf("Unable to move pile")
  }
  if len(pile1) !=0 {
    t.Errorf("Pile has %d but should have 0 card",len(pile1))
  }
  if len(pile2) !=2 {
    t.Errorf("Pile has %d but should have 2 card",len(pile2))
  }
}
