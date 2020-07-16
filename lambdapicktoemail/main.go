package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seth127/funBook/fbutils"
	"github.com/seth127/funBook/funbook"
	"log"
)

///////////// RUN THIS TO COMPILE IT AND SEND TO LAMBDA /////////
//cd lambdapicktoemail
//GOOS=linux GOARCH=amd64 go build -o main main.go
//zip funBookLambdaSes.zip main
//aws s3 cp funBookLambdaSes.zip s3://all-tools-in-the-shadow-of-the-moon/funbook/lambda/funBookLambdaSes.zip
// # then you _might_ need to re-enter that s3 path in the Lambda code section to get it to re-build (not sure)
/////////////////////////////////////////////////////////////////


// Picks from fmt.Printf("s3://%s/%s", fbutils.S3_BUCKET, fbutils.S3_OUT_KEY)
// and currently email everyone who is a verified SES recipient
func handleRequest() (string, error) {


	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(fbutils.S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	pick, pickBytes, pickPath := funbook.GetPickS3(s)
	fmt.Printf("\n---- %d ----\n%s\npickPath: %s\n--------\n", pick, pickBytes, pickPath)

	recipients := funbook.GetValidSesAddresses(s) // get all verified emails from SES

	// build email
	subject := fmt.Sprintf("paragraph %d", pick)
	//body := fmt.Sprintf("\r\n---- %d ----\r\n%s\r\n", pick, pickBytes) // the newlines aren't working. Revisit later.
	body := string(pickBytes)

	funbook.SendSESEmail(s, fbutils.ALL_TOOLS_EMAIL, recipients, subject, body)

	outMsg := fmt.Sprintf("\n---- %d ----\n%s\nEmail sent to %d addresses\n", pick, pickBytes, len(recipients))

	return outMsg, nil
}

func main() {
	lambda.Start(handleRequest)
}
