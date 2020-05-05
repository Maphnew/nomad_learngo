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
