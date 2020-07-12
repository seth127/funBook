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
)

func main() {
	// pick random line
	maxLines := 10000
	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(maxLines)
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
		}

		if pc == pick {
			fmt.Printf("%d\t%v\n", pc, t)
		} else if pc > pick {
			break
		}
	}



	//log.Println("Number of bytes copied to STDOUT:", n)
}
