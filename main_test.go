package main

import "testing"

func TestValidFile(t *testing.T) {
	err := getWordList("mobydick.txt")
	if err != nil {
		t.Error("Application should show the list of 20 words")
	}
}

func TestInvalidFile(t *testing.T) {
	err := getWordList("mobydick.zip")
	if err == nil {
		t.Error("Application should refuse an invalid file: ", err)
	}
}

func TestBinaryFile(t *testing.T) {
	err := getWordList("/bin/cat")
	if err != nil {
		t.Error("Application should show the result of a binary file: ", err)
	}
}
