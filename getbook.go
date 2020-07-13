package main

import (
	"os"
	"bufio"
	"log"
	"time"
	"fmt"
	"strings"

	"net/http"
	"math/rand"

	"github.com/seth127/funBook/funbook"
)

const maxParagraphs = 10

const bookUrl = "https://www.gutenberg.org/files/2701/2701-h/2701-h.htm"
const outDir = "text/MobyDick/"

func PickRand(n int) int {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(n)
	return pick
}

func padNumberWithZero(value uint32) string {
	return fmt.Sprintf("%05d", value)
}

func WriteParagraph(s string, n int) {

	filename := outDir + padNumberWithZero(uint32(n))

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if _, ok := err.(*os.PathError); ok {
		f, err = os.Create(filename)
	}

	checkPanic(err)

	defer f.Close()

	_, err = f.WriteString(s + "\n")
	checkPanic(err)
}


func checkPanic(e error) {
	if e != nil {

		if naw, ok := e.(*os.PathError); ok {
			fmt.Printf("\n-- e.(*os.PathError) -- naw: %v -- ok: %v --%v\n", naw, ok, e.(*os.PathError))
			// e.Name wasn't found
		}

		fmt.Printf("\n%T\t%v\n", e, e)
		panic(e)
	}
}

func main() {
	// pick := PickRand(maxParagraphs)

	// Make HTTP GET request
	response, err := http.Get(bookUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines) // Set up the split function.

	pc := 0
	for scanner.Scan() {

		t := scanner.Text()

		if strings.Contains(t, "<p>") {
			pc++
			if pc > maxParagraphs {
				fmt.Printf("All done. Wrote %d paragraphs.\n", pc-1)
				break
			}
			fmt.Printf("----- %d -----\n", pc)
		}

		pb, pt := funbook.ParseHtml(t)
		if pb {
			//fmt.Printf(pt)
			WriteParagraph(pt, pc)
		}

	}



	//log.Println("Number of bytes copied to STDOUT:", n)
}

