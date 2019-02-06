package file

import "testing"

func getFilePointer(t *testing.T) *fileApp {
	fileApp, err := New()
	if err != nil {
		t.Error("Error get a pointer from fileApp")
	}

	return fileApp
}

func TestEmptyFile(t *testing.T) {
	file := getFilePointer(t)
	_, err := file.GetFile("")
	if err == nil {
		t.Error("Cannot continue with empty file")
	}
}

func TestInvalidFile(t *testing.T) {
	file := getFilePointer(t)
	_, err := file.GetFile("this.txt")
	if err == nil {
		t.Error("Cannot continue with invalid file")
	}
}

func TestValidFile(t *testing.T) {
	file := getFilePointer(t)
	_, err := file.GetFile("../../mobydick.txt")
	if err != nil {
		t.Error("Problem with valid file: ", err)
	}
}

func TestValidWords(t *testing.T) {
	file := getFilePointer(t)
	w, err := file.GetWords("../../mobydick.txt")
	if err != nil {
		t.Error("Problem getting words from a valid file: ", err)
	}

	if len(w) == 0 {
		t.Fail()
	}
}
