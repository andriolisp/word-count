package words

import "testing"

func TestRandomWords(t *testing.T) {
	wa := New([]string{"a", "b", "c", "d", "a", "a", "b", "b"})
	if len(wa) > 3 {
		t.Fail()
	}
}

func TestValidateWords(t *testing.T) {
	wa := New([]string{"a", "b", "c", "d", "a", "a", "b", "b"})
	if wa[0].Quantity != 3 && wa[0].Word == "a" {
		t.Fail()
	}

	if wa[1].Quantity != 3 && wa[0].Word == "b" {
		t.Fail()
	}
}

func TestSortWords(t *testing.T) {
	wa := New([]string{"b", "a", "b", "c", "d", "a", "a", "b", "b"})
	wa.SortByQuantity()
	if wa[0].Quantity != 4 && wa[0].Word == "b" {
		t.Fail()
	}
}

func TestSliceWords(t *testing.T) {
	wa := New([]string{"b", "a", "b", "c", "d", "a", "a", "b", "b"})
	wa.SortByQuantity()

	result := wa.Slice(0, 3)
	if len(result) != 3 {
		t.Error("Array need 3 items")
	}

	if result[0].Quantity != 4 && result[0].Word == "b" {
		t.Error("First index should be b with 4")
	}

	if result[1].Quantity != 3 && result[1].Word == "a" {
		t.Error("First index should be a with 3")
	}

	if result[2].Quantity != 1 && result[2].Word == "c" {
		t.Error("First index should be c with 1")
	}
}
