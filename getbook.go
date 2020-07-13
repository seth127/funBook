package main

import (
	"bufio"
	"log"
	"time"

	"strings"

	//"io"
	//"os"

	"net/http"
	"fmt"

	"math/rand"

	//"funbook"
	"regexp"
)

// regex

var parseDict = map[string]string {
	"&rsquo;" : "'",
	"&mdash;" : "--",
	"&ldquo" : `"`,
	"&rdquo" : `"`,
}

func ParseHtml(s string) (bool, string) {

	// trim whitespace
	ms := regexp.MustCompile(` +`)
	s = ms.ReplaceAllString(s, " ")

	// parse from the dict
	for old, new := range parseDict {
		s = strings.ReplaceAll(s, old, new)
	}

	// replace paragraph
	mp := regexp.MustCompile(`<\/?p>`)
	s = mp.ReplaceAllString(s, "")

	// if there are no letters return false
	ml := regexp.MustCompile("[A-Za-z]")
	b :=  ml.MatchString(s)
	return b, s
}



func main() {
	// pick random line
	maxLines := 99
	rand.Seed(time.Now().UnixNano())
	pick := maxLines // rand.Intn(maxLines)
	fmt.Println(pick)



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
		}

		pb, pt := ParseHtml(t)
		if pb {
			fmt.Printf(pt)
		}

		if pc == pick {
			//pt := strings.ReplaceAll(t, "<p>", "")

			fmt.Println("naw")

		} else if pc > pick {
			break
		}
	}



	//log.Println("Number of bytes copied to STDOUT:", n)
}

