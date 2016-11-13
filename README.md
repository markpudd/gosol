# gosol - Solitaire in Go

This is a simple ASCII based Solitaire implementation written to learn go.

Screenshot :-

`T   Y       A   B   C   D`
`XXX ooo     XXX XXX XXX XXX`
``
`1   2   3   4   5   6   7`
`8SB XXX XXX XXX XXX XXX XXX`
`    KCB XXX XXX XXX XXX XXX`
`        9HR XXX XXX XXX XXX`
`            8DR XXX XXX XXX`
`                KSB XXX XXX`
`                    ASB XXX`
`                        4CB`

To play:-

`  TT - Take a card from deck`
`  QQ - Quit`
`  NN - New game`
`  Y{A-D}     - Move turned card from turned deck to end pile`
`  Y{1-7}     - Move turned card from turned deck to play pile`
`  {1-7}{A-D} - Move play pile to end pile`
`  {1-7}{1-7} - Move play pile to other play pile`
