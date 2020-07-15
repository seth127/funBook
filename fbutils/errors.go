package fbutils

import (
	"os"
	"fmt"
	"strings"
)

// If the error is not nil, first print the type, then panic
func PanicNil(e error) {
	if e != nil {

		fmt.Printf("\n%T\t%v\n", e, e)
		panic(e)
	}
}

// if the error is either nil or *os.PathError do nothing
// otherwise print the type, then panic
func PanicNonPathError(e error) {
	if e != nil {
		if _, ok := e.(*os.PathError); !ok {
			fmt.Printf("\n%T\t%v\n", e, e)
			panic(e)
		}
	}
}

// if the error is either nil or `*s3err.RequestFailure -- NoSuchKey` do nothing
// otherwise print the type, then panic
func PanicNonS3Missing(e error) {
	if e != nil && !CheckS3MissingError(e) {
		fmt.Printf("\n%T\t%v\n", e, e)
		panic(e)
	}
}

// Check if an error is because the S3 key is missing.
// This feels hacky, but I can't access s3err because it's private so I can't check the type.
func CheckS3MissingError(e error) bool {
	if e != nil {
		return strings.Contains(e.Error(), "*s3err.RequestFailure -- NoSuchKey: The specified key does not exist.")
	}
	return false
}


