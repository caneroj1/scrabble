package main

import (
  "fmt"
  "os"
  "strconv"
  "github.com/caneroj1/scrabble/go/player"
  "github.com/caneroj1/scrabble/go/reader"
  "github.com/caneroj1/scrabble/go/game"
)

func main() {
  fmt.Println("Let's Go Scrabble!")
  fmt.Print("How many players? ")

  numPlayers, err := strconv.ParseInt(reader.GetString(os.Stdin), 0, 0)

  if err != nil {
    fmt.Printf("There was an issue with getting the number of players: %s\n", err)
    os.Exit(1)
  }

  if numPlayers < 2 {
    fmt.Println("Please play with more people than that.")
    os.Exit(1)
  }

  players := player.CreatePlayers(int(numPlayers))
  g := game.Game { Players: players }
  g.ShufflePlayers()
  g.Play()
  g.Finish()
}
