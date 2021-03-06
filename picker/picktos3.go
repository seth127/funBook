package main

import (
	"fmt"
	"log"

	"github.com/seth127/funBook/fbutils"
	"github.com/seth127/funBook/funbook"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)


// Picks from fmt.Printf("s3://%s/%s", fbutils.S3_BUCKET, fbutils.S3_OUT_KEY)
// and currently write back to fmt.Printf("s3://%s/%s", fbutils.S3_BUCKET, fbutils.S3_MD_TEST_KEY)
// but the plan is to use the first part of this as the lambda function that picks from s3
// and then compose an email from the byte slice that's returned from funbook.GetPickS3()
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
	funbook.AddPickToS3(s, pickBytes, pickPath)

}

