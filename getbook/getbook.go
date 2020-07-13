package main

import (
	"bufio"
	"fmt"
	"strings"

	"net/http"

	"github.com/seth127/funBook/funbook"
	"github.com/seth127/funBook/fbutils"
)


func main() {

	// Make HTTP GET request
	response, err := http.Get(funbook.BookUrl)
	fbutils.CheckPanic(err)

	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines) // Set up the split function.

	pc := 0
	for scanner.Scan() {

		t := scanner.Text()

		if strings.Contains(t, "<p>") {
			pc++
			if pc > funbook.MaxParagraphs {
				fmt.Printf("All done. Wrote %d paragraphs.\n", pc-1)
				break
			}
			fmt.Printf("----- %d -----\n", pc)
		}

		pb, pt := funbook.ParseHtml(t)
		if pb {
			//fmt.Printf(pt)
			funbook.WriteParagraph(pt, pc, funbook.OutDir)
		}

	}

}

