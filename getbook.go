package main

import (
	"bufio"
	"fmt"
	"strings"

	"net/http"

	"github.com/seth127/funBook/funbook"
	"github.com/seth127/funBook/fbutils"
)

const maxParagraphs = 100

const bookUrl = "https://www.gutenberg.org/files/2701/2701-h/2701-h.htm"
const outDir = "text/MobyDick/"


func main() {
	// pick := PickRand(maxParagraphs)

	// Make HTTP GET request
	response, err := http.Get(bookUrl)
	fbutils.CheckPanic(err)

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
			funbook.WriteParagraph(pt, pc, outDir)
		}

	}

}

