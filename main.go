package main

import (
	"fmt"
	"os"

	"github.com/andriolisp/word-count/app/file"
	"github.com/andriolisp/word-count/app/words"
)

func main() {
	err := getWordList(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func getWordList(filename string) error {
	f, err := file.New()
	if err != nil {
		return err
	}

	//Open the file and return an word list
	wordList, err := f.GetWords(filename)
	if err != nil {
		return err
	}

	//Transform the list of words in a struct with the word + quantity
	w := words.New(wordList)

	//Sort all the list by the quantity
	w.SortByQuantity()

	for _, v := range w.Slice(0, 20) {
		quantity := fmt.Sprintf("%v", v.Quantity)
		spaces := 5 - len(quantity)
		text := fmt.Sprintf("%s %s", quantity, v.Word)

		//Add spaces on the left side to better presentation on the console
		for spaces > 0 {
			text = " " + text
			spaces--
		}

		fmt.Println(text)
	}

	return nil
}
