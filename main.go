package main

import "fmt"
import "github.com/foxreymann/shares/bank"
import "math/rand"

func main() {
  ledger := bank.Ledger
	fmt.Println(ledger)

  bank.Withdraw("A", "Fox", rand.Intn(10))
}
