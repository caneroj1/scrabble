package game

import (
  "time"
  "fmt"
  "os"
  "sort"
  "math/rand"
  "strconv"
  "github.com/caneroj1/scrabble/go/player"
  "github.com/caneroj1/scrabble/go/reader"
)

type Game struct {
  Players player.Players
}

func (g Game) ShufflePlayers() {
  rnd := rand.New(rand.NewSource(time.Now().Unix()))
  for i := len(g.Players) - 1; i > 0; i-- {
    var tmp int = rnd.Intn(i + 1)
    g.Players[i], g.Players[tmp] = g.Players[tmp], g.Players[i]
  }
}

func (g Game) Finish() {
  sort.Sort(g.Players)
  fmt.Println("Finished!")
  fmt.Printf("%s won!\n", g.Players[0].Name)
  for _, p := range g.Players {
    fmt.Printf("%s scored %d points\n", p.Name, p.Score)
  }
}

func (g Game) Play() {
  var input string = ""
  var err error = nil
  var score int64 = 0
  i := 0

  fmt.Printf("%s goes first!\n", g.Players[i].Name)
  for {
    fmt.Printf("%s scored: ", g.Players[i].Name)

    input = reader.GetString(os.Stdin)
    if input == "done" {
      break
    }

    score, err = strconv.ParseInt(input, 0, 0)

    if err != nil {
        fmt.Printf("Uh oh! There was a problem with that input: %s\n", err)
        os.Exit(1)
    }

    g.Players[i].IncreaseScore(int(score))

    i = (i + 1) % len(g.Players)
  }
}
