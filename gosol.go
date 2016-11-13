package main

import (
  "fmt"
  "github.com/markpudd/gosol/gosol"
)



func main() {
  var f,t byte
  g := common.NewGame()
  g.ShowBoard()
  fmt.Scanf("%c%c\n",&f,&t)
  for g.Turn(f,t) {
    fmt.Scanf("%c%c\n",&f,&t)
  }


}
