package main

import "github.com/foxreymann/shares/bank"
import "fmt"
import "time"
import "math/rand"

func main() {
  // init random number generator
  rand := rand.New(rand.NewSource(time.Now().UnixNano()))

  _, err := bank.Withdraw("A", "fox", rand.Intn(2))

  if err != nil {
    fmt.Println(err)
  }

  _, err = bank.Withdraw("A", "paul", rand.Intn(2))

  if err != nil {
    fmt.Println(err)
  }

  _, err = bank.Transfer("A", "paul", "fox")

  if err != nil {
    fmt.Println(err)
  }
}
