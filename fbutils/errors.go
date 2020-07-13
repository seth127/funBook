package fbutils

import (
	"os"
	"fmt"
)

func CheckPanic(e error) {
	if e != nil {

		if naw, ok := e.(*os.PathError); ok {
			fmt.Printf("\n-- e.(*os.PathError) -- naw: %v -- ok: %v --%v\n", naw, ok, e.(*os.PathError))
			// e.Name wasn't found
		}

		fmt.Printf("\n%T\t%v\n", e, e)
		panic(e)
	}
}
