package main

import (
	"fmt"
	"github.com/seth127/funBook/funbook"
)

// This will eventually become a package function call
// and more specifically, the call that will be made
// in the lambda.
func main() {
	pick, pickString, _ := funbook.GetPickLocal()

	fmt.Printf("\n ------ %d -------\n", pick)
	fmt.Println(pickString)

	//pick2, _, pickPath := funbook.GetPickLocal()
	//
	//fmt.Printf("\n ------ %d -------\n", pick2)
	//pathSlice := strings.Split(pickPath, "/")
	//fmt.Println(pathSlice[len(pathSlice)-1])
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Format("20060102-150405"))
	//
	//// build destination path
	//timeStamp := time.Now().Format("20060102-150405")
	//
	//destFile := pathSlice[len(pathSlice)-1] + "-" + timeStamp
	//
	//destPath := filepath.Join(fbutils.S3_MD_TEST_KEY, destFile)
	//fmt.Printf("s3://%s/%s\n", fbutils.S3_BUCKET, destPath)
}