package main

import (
	"fmt"

	"github.com/maphnew/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Maphnew")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
