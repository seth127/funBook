package main

import (
	"bufio"
	"log"
	"time"

	"strings"

	//"io/ioutil"
	"os"

	"net/http"
	"fmt"

	"math/rand"

	"github.com/seth127/funBook/funbook"
)

const maxParagraphs = 99
const outDir = "text/MobyDick/"

func PickRand(n int) int {
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(n)
	return pick
}

func WriteParagraph(s string, n int) {

	filename := outDir + string(n)

	f, err := os.Open(filename)

	if os.IsNotExist(err) {
		f, err = os.Create(filename)
	}
	checkPanic(err)

	defer f.Close()

	_, err = f.WriteString(s)
	checkPanic(err)
}


func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// pick := PickRand(maxParagraphs)

	// Make HTTP GET request
	response, err := http.Get("https://www.gutenberg.org/files/2701/2701-h/2701-h.htm")
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
			fmt.Printf("\n----- %d -----\n", pc)
			if pc > maxParagraphs {
				break
			}
		}

		pb, pt := funbook.ParseHtml(t)
		if pb {
			fmt.Printf(pt)

			if pc == maxParagraphs {
				WriteParagraph(pt, pc)
			}
		}

	}



	//log.Println("Number of bytes copied to STDOUT:", n)
}

