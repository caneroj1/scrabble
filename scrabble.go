package main

import (
  "fmt"
  "os"
  "strconv"
  "github.com/caneroj1/scrabble/go/player"
  "github.com/caneroj1/scrabble/go/reader"
)

func main() {
  fmt.Println("Let's Go Scrabble!")
  fmt.Print("How many players? ")

  numPlayers, err := strconv.ParseInt(reader.GetString(os.Stdin), 0, 0)

  if err != nil {
    fmt.Printf("There was an issue with getting the number of players: %s\n", err)
    os.Exit(1)
  }

  if numPlayers < 0 {
    fmt.Println("The number of players cannot be negative.")
    os.Exit(1)
  }

  players := player.CreatePlayers(int(numPlayers))

  for i, v := range players {
    fmt.Printf("Player %d: %s\n", i + 1, v.Name)
  }
}
