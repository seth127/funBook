package main

import (
	"bufio"
	"log"

	//"strings"

	//"io"
	//"os"

	"net/http"
	"fmt"
)

func main() {
	// Make HTTP GET request
	response, err := http.Get("https://www.devdungeon.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	//n, err := io.Copy(os.Stdout, response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//buf := make([]byte, 8192)
	//for i := 1; i <= 10; i++ {
	//	n, error := response.Body.Read(buf)
	//
	//	if error == io.EOF {
	//		break
	//	} else if error != nil {
	//		log.Fatal(error)
	//	}
	//
	//	fmt.Printf("%d\ttype: %T\tthing: %v\t%v\n", i, n, n, string(buf[:n])[:10])
	//
	//}

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines) // Set up the split function.

	//buf := make([]byte, 8192)
	i := 0
	for scanner.Scan() {
		i++
		t := scanner.Text()

		fmt.Printf("%d\ttype: %T\tthing: %v\n", i, t, t)

		if i > 10 {
			break
		}
	}



	//log.Println("Number of bytes copied to STDOUT:", n)
}
