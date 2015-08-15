package player

import (
  "fmt"
  "github.com/caneroj1/scrabble/go/reader"
  "os"
)

type Player struct {
  Name string
  Score int
}

type Players []Player

func CreatePlayers(n int) Players {
  players := make(Players, n)

  for i := 0; i < n; i++ {
    fmt.Printf("Name of player %d: ", i + 1)
    players[i] = Player{ reader.GetString(os.Stdin), 0 }
  }

  return players
}

func (p *Player) IncreaseScore(s int) {
  p.Score += s
}

func (p Players) Len() int {
  return len(p)
}

func (p Players) Less(i, j int) bool {
  return !(p[i].Score < p[j].Score)
}

func (p Players) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
}
