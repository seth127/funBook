package main

import (
	"fmt"
	"os"

	"github.com/seth127/funBook/funbook"
	"github.com/seth127/funBook/fbutils"
)


func main() {
	pick := fbutils.PickRand(funbook.MaxParagraphs)
	pickString, pickErr := funbook.ReadParagraph(funbook.OutDir, pick)
	fbutils.CheckNonMissing(pickErr)
	_, missingErr := pickErr.(*os.PathError)

	for missingErr {
		//fmt.Printf("\n ------ %d ------- REDO\n", pick)
		pick = fbutils.PickRand(funbook.MaxParagraphs)
		pickString, pickErr = funbook.ReadParagraph(funbook.OutDir, pick)
		fbutils.CheckNonMissing(pickErr)
		_, missingErr = pickErr.(*os.PathError)
	}

	fmt.Printf("\n ------ %d -------\n", pick)
	fmt.Println(pickString)
}