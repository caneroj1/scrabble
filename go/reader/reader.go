package reader

import (
  "bufio"
  "io"
  "os"
  "fmt"
)

func GetString(r io.Reader) string {
  scanner := bufio.NewScanner(r)
  scanner.Scan()

  if scanner.Err() != nil {
    fmt.Printf("Something went wrong: %s\n", scanner.Err())
    os.Exit(1)
  }

  return scanner.Text();
}
