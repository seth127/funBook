package main

import (
	"fmt"
	"github.com/seth127/funBook/funbook"
)

// Picks from local ./text/MobyDick/ folder (which is .gitignore'd)
// and write to console. Only for testing.
func main() {
	pick, pickString, _ := funbook.GetPickLocal()

	fmt.Printf("\n ------ %d -------\n", pick)
	fmt.Println(pickString)

}