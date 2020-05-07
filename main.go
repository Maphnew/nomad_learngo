package main

import (
	"fmt"

	"github.com/maphnew/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Maphnew")
	fmt.Println(account)
}
