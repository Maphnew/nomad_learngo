package main

import (
	"fmt"

	"github.com/maphnew/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println("err:", err)
	}

	hello, _ := dictionary.Search("hello")
	fmt.Println("found:", word, "definition:", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println("err2:", err2)
	}
}
