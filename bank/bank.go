package bank

import "fmt"
import "errors"

type Share struct {
  Type string
  bookValue int
  Owner string
}

var Ledger []Share

var lottery map[string]int

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

 lottery = make(map[string]int)
 lottery["fox"] = 0
 lottery["paul"] = 1
}

func Withdraw(shareType string, newOwner string, lotteryNumber int) (*Share, error) {
  fmt.Println(lotteryNumber)

  _, isOwnerInLottery := lottery[newOwner]

  if !isOwnerInLottery {
    return nil, errors.New("Owner not in the lottery")
  }

  // check lottery number
  if lottery[newOwner] != lotteryNumber {
    return nil, errors.New("Owner didn't win the lottery")
  }

  // run Transfer function
  return Transfer(shareType, "bank", newOwner)
}

func Transfer(shareType string, owner string, newOwner string) (*Share, error) {
  // check if owner has balance

  var shareToTransfer *Share = nil

  for _, share := range Ledger {
    if share.Type == shareType && share.Owner == owner {
      fmt.Println(share)
      shareToTransfer = &share
      break
    }
  }

  fmt.Println(shareToTransfer)


  // make the transfer
  return nil, errors.New("not yet implmenented")
}
