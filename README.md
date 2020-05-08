# Learn Go
## Contents
### 1.0 Theory
- 1.0 Main Package
- 1.1 Packages and Imports
- 1.2 Variables and Constants
- 1.3 Functions part One
- 1.4 Functions part Two
- 1.5 for, range, ...args
- 1.6 If with a Twist
- 1.7 Switch
- 1.8 Pointers!
- 1.9 Arrays and Slices
- 1.10 Maps
- 1.11 Structs

### 2 Bank & Dictionary Projects
- 2.0 Account + NewAccount
- 2.1 Methods part One
- 2.2 Methods part Two
- 2.3 Finishing Up
- 2.4 Dictionary part One
- 2.5 Add Method
- 2.6 Update Delete

### 3 URL Checker & Go Routines
 
- 3.0 hitURL()
- 3.1 Slow URLChecker
- 3.2 Goroutines
- 3.3 Channels
- 3.4 Channels Recap
- 3.5 One more Recap
- 3.6 URLChecker + Go Routines
- 3.7 FAST URLChecker

### 4 Job Scrapper
 
- 4.0 getPages part One
- 4.1 getPages part Two
- 4.2 extractJob part One
- 4.3 extractJob part Two
- 4.4 Writing Jobs
- 4.5 Channels Time
- 4.6 More Channels Baby
- 4.7 Recap

### 5 Web Server with Echo
 
- 5.0 Setup Part One
- 5.1 Setup Part Two
- 5.2 File Download
- 5.3 Conclusions

### 1.0 Theory
```go
// main.go

package main

// 1.0 Main Package
// 1.1 Packages and Imports
import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	// 1.3 functions part one
	// naked return
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return

	// return len(name), strings.ToUpper(name)
}

func repeatMe(a int, words ...string) {
	// 1.4 functions part two
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	// 1.5 for, range, ...args
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }
	total := 0
	for _, number := range numbers {
		// fmt.Println(number)
		total += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	return total
}

func canIDrink(age int) bool {
	// 1.6 variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func canIDrinkSwitch(age int) bool {
	// 1.7 switch
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

}

func main() {
	// 1.2 Variables and Constants
	const name string = "maph"

	fmt.Println(name)

	var name2 = "maph"
	name2 = "park"
	fmt.Println(name2)

	name3 := "mappppph"
	name3 = "new"
	fmt.Println(name3)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("maph")
	fmt.Println(totalLength, upperName)

	totalLength2, _ := lenAndUpper("maph")
	fmt.Println(totalLength2)

	repeatMe(3, "aa", "bb", "aasc")

	result := superAdd(12, 13, 14, 15)
	fmt.Println(result)
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkSwitch(18))

	// 1.8 Pointers
	// & address
	// * see through
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)

	// 1.9 Arrays and Slices
	names := []string{"a", "b", "c"} // slice, array without length
	names2 := append(names, "d")

	fmt.Println(names, names2)

	// 1.10 Maps
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	fmt.Println(nico["name"], nico["age"])

	for _, value := range nico {
		fmt.Println(value)
	}

	// 1.11 Structs
	favFood := []string{"water", "drink"}
	maph := person{"maphnew", 18, favFood}
	fmt.Println(maph.name, maph.age)

	maph2 := person{name: "maphnew", age: 18, favFood: favFood}
	fmt.Println(maph2.name, maph2.age)

}

type person struct {
	name    string
	age     int
	favFood []string
}

```

### 2 Bank & Dictionary Projects
#### 2.0 Account + NewAccount

- Constructor(http://www.golangpatterns.info/object-oriented/constructors)
- Uppercase for Export 
- Public/Private
- Golint - Comments for Export someting
- Pointer, Address

```go
// main.go

package main

import (
	"fmt"
	"github.com/maphnew/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Maphnew")
	fmt.Println(account)
}
```
```go
// accounts/accounts.go

package accounts

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

```
#### 2.1 Methods part One
```go
// main.go
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

```
```go
// accounts/accounts.go
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

```
#### 2.2 Methods part Two
#### 2.3 Finishing Up
#### 2.4 Dictionary part One
#### 2.5 Add Method
#### 2.6 Update Delete

### 3 URL Checker & Go Routines
 
#### 3.0 hitURL()
#### 3.1 Slow URLChecker
#### 3.2 Goroutines
#### 3.3 Channels
#### 3.4 Channels Recap
#### 3.5 One more Recap
#### 3.6 URLChecker + Go Routines
#### 3.7 FAST URLChecker

### 4 Job Scrapper
 
#### 4.0 getPages part One
#### 4.1 getPages part Two
#### 4.2 extractJob part One
#### 4.3 extractJob part Two
#### 4.4 Writing Jobs
#### 4.5 Channels Time
#### 4.6 More Channels Baby
#### 4.7 Recap

### 5 Web Server with Echo
 
#### 5.0 Setup Part One
#### 5.1 Setup Part Two
#### 5.2 File Download
#### 5.3 Conclusions