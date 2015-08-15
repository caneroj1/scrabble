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

func CreatePlayers(n int) []Player {
  players := make([]Player, n)

  for i := 0; i < n; i++ {
    fmt.Printf("Name of player %d: ", i + 1)
    players[i] = Player{ reader.GetString(os.Stdin), 0 }
  }

  return players
}
