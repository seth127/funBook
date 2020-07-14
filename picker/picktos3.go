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

// This all needs to go in an aws module


//////////

func main() {
	//pick, pickPath := funbook.GetPickS3()

	//fmt.Printf("\n ------ %d -------\n", pick)

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(fbutils.S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	/////// the pick part will go away
	//pick := fbutils.PickRand(fbutils.MaxParagraphs)
	//
	//pickPath := filepath.Join(fbutils.S3_OUT_KEY, fbutils.PadNumberWithZero(uint32(pick)))
	//fmt.Printf("Trying to pull s3://%s/%s\n", fbutils.S3_BUCKET, pickPath)
	////////
	//
	//err = funbook.ReadFileFromS3(s, fbutils.S3_BUCKET, pickPath)

	pick, destination, err := funbook.GetPickS3(s)
	if err != nil {
		log.Fatal(err)
	}

	//pick := fbutils.PickRand(fbutils.MaxParagraphs)
	//pickPath := filepath.Join(fbutils.S3_OUT_KEY, fbutils.PadNumberWithZero(uint32(pick)))
	////pickString, pickPath,
	//pickErr := funbook.ReadFileFromS3(s, fbutils.S3_BUCKET, pickPath)
	//if pickErr != nil {
	//	log.Fatal(pickErr)
	//}


	//fmt.Printf("Successfully wrote to s3://%s/%s\n", fbutils.S3_BUCKET, destPath)
	fmt.Printf("picktos3 -- %v -- %v \n", pick, destination)
	fmt.Printf("picktos3 successfully read from s3://%s/%s\n", fbutils.S3_BUCKET, destination)

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