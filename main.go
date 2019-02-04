package main

import "github.com/foxreymann/shares/bank"
import "time"
import "math/rand"

func main() {
  // init random number generator
  rand := rand.New(rand.NewSource(time.Now().UnixNano()))

  bank.Withdraw("A", "fox", rand.Intn(4))
}
