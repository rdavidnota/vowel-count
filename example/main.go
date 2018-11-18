package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	InputPath  string
	OutputPath string
	InputFile  []byte
	OutputFile []byte
	InputText  string
	RegEx       = "[a-zA-Z]+"
	RegExVowels = "[aeiouAEIOU]+"
	MapWords    map[string]int
)

func main() {
	InputPath = "input.txt"
	OutputPath = "output.txt"

	LoadFile(InputPath)
	CounterByWordFrequency()
	SortByKey()
}

func SortByKey() {
	names := make([]string, 0, len(MapWords))

	for name := range MapWords {
		names = append(names, name)
	}

	sort.Strings(names)

	var auxiliar string

	for _, name := range names {
		var line string
		line = fmt.Sprintf("%s: %d \n", name, MapWords[name])

		auxiliar += line
	}

	SaveToFile(OutputPath, auxiliar)
}

func CounterByWordFrequency(){
	var words []string
	MapWords = make(map[string]int)

	words = SplitWords()

	for _, word := range words {

		if _, exits := MapWords[word]; exits{
			MapWords[word] += 1
		} else{
			MapWords[word] = 1
		}
	}
}

func CounterByNumbersVowels() {
	var words []string
	MapWords = make(map[string]int)

	words = SplitWords()

	for _, word := range words {

		if _, exits := MapWords[word]; exits{
			MapWords[word] += VowelCounter(word)
		} else{
			MapWords[word] = VowelCounter(word)
		}
	}
}

func VowelCounter(word string) int {
	var count int
	count = 0

	for _, char := range word {
		var auxiliar string
		auxiliar = string(char)

		match, _ := regexp.MatchString(RegExVowels, auxiliar)

		if match {
			count += 1
		}
	}

	return count
}

func SplitWords() []string {

	var words []string

	r, _ := regexp.Compile(RegEx)
	words = r.FindAllString(InputText, -1)

	var counter int
	counter = 0

	for counter < len(words) {
		words[counter] = strings.ToLower(words[counter])
		counter += 1
	}

	return words
}

func SaveToFile(path string , word string) {
	file, err := os.OpenFile(path,os.O_APPEND|os.O_WRONLY,0600)
	check(err)

	defer file.Close()

	_, err = file.WriteString(word)
	check(err)

	file.Sync()
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFile(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)
	InputFile = dat

	InputText = string(InputFile)
}