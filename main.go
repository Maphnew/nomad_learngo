package main

import (
	"fmt"

	"github.com/maphnew/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Maphnew")
	account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
