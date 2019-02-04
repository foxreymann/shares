package bank

import "fmt"

type Share struct {
  Type string
  bookValue int
  Owner string
}

var Ledger []Share

type LotteryRule struct {
  Owner string
  winningNumber int
}

var LotteryRules []LotteryRule

func init() {
  Ledger = []Share{
    {"A", 10,"bank"},
    {"A", 10,"bank"},
    {"A", 10,"bank"},
    {"B", 5,"bank"},
    {"B", 5,"bank"},
    {"B", 5,"bank"},
    {"C", 3,"bank"},
    {"C", 3,"bank"},
    {"C", 3,"bank"},
  }

  LotteryRules = []LotteryRule{
    {"Fox", 1},
    {"Paul", 2},
  }
}

func Withdraw(shareType string, newOwner string, lotteryNumber int) bool {
  fmt.Println(lotteryNumber)
  // check lottery number
  // run Transfer function
  return Transfer(shareType, newOwner)
}

func Transfer(shareType string, newOwner string) bool {
  return true
}
