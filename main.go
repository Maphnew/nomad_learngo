package main

import (
	"fmt"

	"github.com/maphnew/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
