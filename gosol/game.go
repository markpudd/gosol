package gosol

import (
  "fmt"
)

type Game struct {
  deck *Deck
  turnDeck *Deck
  playPiles [7]Pile
  finishPiles [4]Pile
}

func (g *Game) Init() {
  g.deck = NewDeck(52)
  g.turnDeck = NewDeck(0)
  g.deck.Shuffle()
  for i := 0; i < 7; i++ {
    g.playPiles[i]= CreatePile()
  }
  for i := 0; i < 4; i++ {
    g.finishPiles[i]= CreatePile()
  }
  for i := 0; i < 7; i++ {
    for y := 0; y <= i; y++ {
      g.playPiles[i]= append(g.playPiles[i],g.deck.TakeCard());
      if y == i {
        g.playPiles[i][y].FlipCard()
      }
    }
  }
}

func (g *Game) MoveFromTo(f, t byte) {
  g.playPiles[f].MoveStack(&g.playPiles[t])
}

func (g *Game) MoveToEnd(f, t byte) {
  card := g.playPiles[f].LastCard()
  if card != nil {
    if g.finishPiles[t].AddEndCard(card) {
       g.playPiles[f].RemoveLastCard()
    }
  }
}
func (g *Game) FromDeck( t byte) {
  card := g.turnDeck.TopCard()
  if t >= 'A' && t <='D' {
    if g.finishPiles[t-'A'].AddEndCard(card) {
       g.turnDeck.TakeCard()
    }
  } else if t >= '1' && t <='7' {
    if g.playPiles[t-'1'].AddPlayCard(card) {
       g.turnDeck.TakeCard()
    }
  }
}

func (g *Game) TurnTop() {
  card := g.deck.TakeCard()
  if card != nil {
    card.FlipCard()
    g.turnDeck.AddCard(card)
  } else {
    g.turnDeck.FlipDeckTo(g.deck)
  }
}

func (g *Game) CheckWin() bool {
  for i:=0;i<4;i++ {
    if g.finishPiles[i].LastCard() == nil || g.finishPiles[i].LastCard().card >> 2 != 12 {
      return false
    }
  }
  return true
}

func (g *Game) Turn(f, t byte) bool{
  fmt.Printf("%c %c\n",f,t)
  switch f {
    case 'T':
      g.TurnTop()
    case 'Q':
      return false
    case 'N':
      g.Init()
    case 'Y':
      g.FromDeck(t)
    default:
      if t >= 'A' && t <='D' && f >= '1' && f <='7'{
          g.MoveToEnd(f-'1', t-'A')
      } else if t >= '1' && t <='7' && f >= '1' && f <='7' {
          g.MoveFromTo(f-'1', t-'1')
      }
  }
  g.ShowBoard()
  if g.CheckWin() {
      fmt.Println("\x1B[32mYou Win!!!!!\033[0m")
      // finish game
      return false;
  }
  return true
}


func NewGame() *Game {
  g := new(Game)
  g.Init()
  return g;
}


func (g *Game) ShowBoard() {
  fmt.Println("\x1B[34mT   Y       A   B   C   D\033[0m")

  if g.deck.TopCard() != nil {
    fmt.Print(g.deck.TopCard().String())
  } else {
      fmt.Print("ooo")
  }
  fmt.Print(" ")
  if g.turnDeck.TopCard() != nil {
    fmt.Print(g.turnDeck.TopCard().String())
  } else {
      fmt.Print("ooo")
  }
  fmt.Print("     ")

  for i := 0; i < 4; i++ {
    if len(g.finishPiles[i]) > 0 {
      fmt.Printf("%s ",g.finishPiles[i][len(g.finishPiles[i])-1].String())
    } else {
        fmt.Print("XXX ")
    }
  }
  fmt.Println()
  fmt.Println()
  fmt.Println("\x1B[34m1   2   3   4   5   6   7\033[0m")
  pilesDone := false
  depth :=0
  for  !pilesDone {
    pilesDone = true
    for i := 0; i < 7; i++ {
        if len(g.playPiles[i]) > depth {
            fmt.Printf("%s ",g.playPiles[i][depth].String())
            pilesDone = false
        } else {
            fmt.Print("    ")
        }
	   }
     depth = depth+1
     fmt.Println()
	}
}
