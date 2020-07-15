package main

import (
	"fmt"
	"path/filepath"

	//"fmt"
	"github.com/seth127/funBook/fbutils"
	"github.com/seth127/funBook/funbook"
	"strings"
	"time"

	//"strings"
	//"time"

	"log"
	//"path/filepath"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
)


func main() {

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(fbutils.S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	pick, pickBytes, pickPath := funbook.GetPickS3(s)

	//fmt.Printf("picktos3 successfully read from %s\n", pickPath)
	fmt.Printf("\n---- %d ----\n%s\n", pick, pickBytes)

	/////// This part will change to email at some point
	// build destination path
	pathSlice := strings.Split(pickPath, "/")
	timeStamp := time.Now().Format("20060102-150405")
	destFile :=  fmt.Sprintf("%s-%s", timeStamp, pathSlice[len(pathSlice)-1])
	destPath := filepath.Join(fbutils.S3_MD_TEST_KEY, destFile)

	// Upload
	err = funbook.WriteBufferToS3(s, fbutils.S3_BUCKET, pickBytes, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully wrote to s3://%s/%s\n", fbutils.S3_BUCKET, destPath)


}


func naw() {

	pick, _, pickPath := funbook.GetPickLocal()

	fmt.Printf("\n ------ %d -------\n", pick)




	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(fbutils.S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// build destination path
	pathSlice := strings.Split(pickPath, "/")
	timeStamp := time.Now().Format("20060102-150405")
	destFile :=  fmt.Sprintf("%s-%s", timeStamp, pathSlice[len(pathSlice)-1])
	destPath := filepath.Join(fbutils.S3_MD_TEST_KEY, destFile)

	// Upload
	err = funbook.AddFileToS3(s, fbutils.S3_BUCKET, pickPath, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully wrote to s3://%s/%s\n", fbutils.S3_BUCKET, destPath)
}