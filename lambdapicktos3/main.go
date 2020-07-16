package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seth127/funBook/fbutils"
	"github.com/seth127/funBook/funbook"
)

///////////// RUN THIS TO COMPILE IT AND SEND TO LAMBDA /////////
//GOOS=linux GOARCH=amd64 go build -o main main.go
//zip funBookLambda.zip main
//aws s3 cp funBookLambda.zip s3://all-tools-in-the-shadow-of-the-moon/funbook/lambda/funBookLambda.zip
// # then you _might_ need to re-enter that s3 path in the Lambda code section to get it to re-build (not sure)
/////////////////////////////////////////////////////////////////


// Picks from fmt.Printf("s3://%s/%s", fbutils.S3_BUCKET, fbutils.S3_OUT_KEY)
// and currently write back to fmt.Printf("s3://%s/%s", fbutils.S3_BUCKET, fbutils.S3_MD_TEST_KEY)
// but the plan is to use the first part of this as the lambda function that picks from s3
// and then compose an email from the byte slice that's returned from funbook.GetPickS3()
func handleRequest() (string, error) {

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(fbutils.S3_REGION)})
	if err != nil {
		return "FAILED", err
	}

	pick, pickBytes, pickPath := funbook.GetPickS3(s)

	/////// This part will change to email at some point
	funbook.AddPickToS3(s, pickBytes, pickPath)

	outMsg := fmt.Sprintf("\n---- %d ----\n%s\nSuccessfully wrote to s3://%s/%s/\n", pick, pickBytes, fbutils.S3_BUCKET, fbutils.S3_MD_TEST_KEY)
	return outMsg, nil
}

func main() {
	lambda.Start(handleRequest)
}
