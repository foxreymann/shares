package main

import "fmt"
import "github.com/foxreymann/shares/bank"
import "time"
import "math/rand"

func main() {
  ledger := bank.Ledger
	fmt.Println(ledger)

  // init random number generator
  rand := rand.New(rand.NewSource(time.Now().UnixNano()))

  bank.Withdraw("A", "Fox", rand.Intn(4))
}
