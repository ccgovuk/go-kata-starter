package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlwaysTrue(t *testing.T) {

	assert.True(t, true)
}

func TestKinshipPinkish(t *testing.T) {
	//var dictionary = []string{"pinkish", "polo", "pool", "loop"}
	var dictionary = readDictionary("/usr/share/dict/words")
	result := solve("kinship", dictionary)

	assert.Contains(t, result, "pinkish")
}

func TestPoolLoopPolo(t *testing.T) {

	var dictionary = []string{"pinkish", "polo", "pool", "loop"}

	result := solve("loop", dictionary)

	assert.Equal(t, len(result), 2)
	assert.Contains(t, result, "polo")
	assert.Contains(t, result, "pool")
}

func readDictionary(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	slice := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		slice = append(slice, strings.TrimSuffix(str, "\n"))
	}
	return slice

}

// func readToDisplayUsingFile1(f *os.File){
//     defer f.Close()
//     slice := make([]string,0)

//     reader := bufio.NewReader(f)

//     for{

//     str, err := reader.ReadString('\n')
//     if err == io.EOF{
//         break
//     }

//	    slice = append(slice, str)
//	}
func solve(word string, dictionary []string) []string {

	result := []string{}
	x := runeArray(word)

	for _, wordInDictonary := range dictionary {
		if runeArray(wordInDictonary) == x && wordInDictonary != word {
			result = append(result, wordInDictonary)
		}
	}
	return result
}

func runeArray(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
