package accounts

import "fmt"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount Creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method
// func ([receiver] [type of receiver]) ]method name]([argument] [argment type]) {  }

// Deposit x amount on your account
func (a Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}
