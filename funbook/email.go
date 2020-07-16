package funbook

import (
	"fmt"
	"github.com/seth127/funBook/fbutils"
	"os"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

const (
	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	//ConfigurationSet = "ConfigSet"

	// The HTML body for the email.
	//HtmlBody =  "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
	//	"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
	//	"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	// The character encoding for the email.
	CharSet = "UTF-8"
)

func SendSESEmail(
	s *session.Session,
	Sender string,
	Recipients []string,
	Subject string,
	TextBody string,
	) {

	// Create an SES session.
	svc := ses.New(s)

	// Construct HTML Body
	HtmlBody := fmt.Sprintf("<p>%s</p>", TextBody)

	// Build recipient slice
	var recAwsSlice []*string
	for _, r := range(Recipients) {
		recAwsSlice = append(recAwsSlice, aws.String(r))
	}

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			BccAddresses: recAwsSlice,
			ToAddresses: []*string{
				aws.String(fbutils.ALL_TOOLS_EMAIL),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Printf("Email Sent to %d addresses\n", len(Recipients))
	fmt.Println(result)
}

func GetValidSesAddresses(s *session.Session) []string {

	// Create SES service client
	svc := ses.New(s)

	result, err := svc.ListIdentities(&ses.ListIdentitiesInput{IdentityType: aws.String("EmailAddress")})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var winners []string

	for _, email := range result.Identities {
		var e = []*string{email}

		verified, err := svc.GetIdentityVerificationAttributes(&ses.GetIdentityVerificationAttributesInput{Identities: e})

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, va := range verified.VerificationAttributes {
			if *va.VerificationStatus == "Success" {
				//fmt.Println(*email)
				winners = append(winners, *email)
			}
		}
	}

	return winners
}