//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode"
)

type Total struct {
	count int
	sync.Mutex
}

func analyzeWord(word string, wg *sync.WaitGroup, total *Total) {
	total.Lock()
	defer wg.Done()
	defer total.Unlock()

	wordCount := 0

	for _, letter := range word {
		if unicode.IsLetter(letter) {
			wordCount++
		}
	}

	total.count += wordCount
}

func main() {
	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	totalLetters := Total{}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		line = line[:len(line)-1]
		for _, word := range strings.Split(line, " ") {
			wg.Add(1)
			go analyzeWord(word, &wg, &totalLetters)
		}
	}

	wg.Wait()
	fmt.Printf("Total letters: %v\n", totalLetters.count)

}
