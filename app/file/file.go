package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

//fileApp has all the functions to read the file
type fileApp struct {
	scanner *bufio.Scanner
	words   []string
}

//New return a pointer of the FileApp struct
func New() (*fileApp, error) {
	//Enable the use of the all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	return &fileApp{}, nil
}

//GetFile return a pointer of the os.File struct
func (f *fileApp) GetFile(filename string) (*os.File, error) {
	if len(filename) == 0 {
		return nil, errors.New("filename is mandatory")
	}

	filePath := filename
	if string(filePath[0]) != "/" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		//validate the whole path of the application
		filePath = path.Join(dir, filename)
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s not found", filePath)
	}

	return os.Open(filePath)
}

//GetStream return an array de bytes for the filename
func (f *fileApp) GetWords(filename string) ([]string, error) {
	osFile, err := f.GetFile(filename)
	if err != nil {
		return nil, err
	}
	defer osFile.Close()

	f.words = nil
	f.scanner = bufio.NewScanner(osFile)

	//Divide the full text into words
	f.scanner.Split(bufio.ScanWords)

	done := make(chan bool)
	word := make(chan string)

	//enable a separate thread to get the word
	go f.readWords(word, done)
	//clean the string and add to the array
	go f.feedArray(word)
	<-done

	return f.words, nil
}

func (f *fileApp) readWords(word chan string, done chan bool) {
	for f.scanner.Scan() {
		word <- f.scanner.Text()
	}
	close(word)
	done <- true
}

func (f *fileApp) feedArray(word <-chan string) {
	//Regex to clean anything that is different them A-ZA-z
	reg, _ := regexp.Compile("[^A-Za-z]+")
	for data := range word {
		f.words = append(f.words, strings.ToLower(reg.ReplaceAllString(data, "")))
	}
}
