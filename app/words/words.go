package words

import "sort"

type wordsApp struct {
	Word     string
	Quantity int
}

type wordsInfo []wordsApp

//New will return a pointer of the wordsInfo struct
func New(words []string) wordsInfo {
	q := 1
	w := ""
	wa := make(wordsInfo, 0)

	//Sort all the words on the array
	sort.Strings(words)

	//Group the words on the struct by quantity
	for i := range words {
		if i == 0 {
			w = words[i]
		}

		if w == words[i] {
			q++
		} else {
			wa = append(wa, wordsApp{w, q})
			w = words[i]
			q = 1
		}

		if i == len(words)-1 {
			wa = append(wa, wordsApp{w, q})
		}
	}

	return wa
}

//SortByQuantity will sort the whole struct based on the quantity
func (w wordsInfo) SortByQuantity() {
	//Use the anonymous function to filter by quantity
	sort.Slice(w, func(i, j int) bool {
		return w[i].Quantity > w[j].Quantity
	})
}

//Slice will return a slice of the array based on the size
func (w wordsInfo) Slice(index, size int) wordsInfo {
	//Check if is not out of bounds
	if index+1+size > len(w) {
		return w[index:(len(w) - (index + 1))]
	} else {
		return w[index:size]
	}
}
