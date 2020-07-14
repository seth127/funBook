package funbook

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"net/http"
	"os"
)

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, bucket string, filePath string, destPath string) error {

	// Open the file for use
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(destPath),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}


func ReadFileFromS3(s *session.Session, bucket string, key string) (string, error) {

	destination := fmt.Sprintf("s3://%s/%s", bucket, key)

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(s)


	// Write to a file !!!!!!!!!
	// Create a file to write the S3 Object contents to.
	f, err := os.Create("TESTER.txt")
	if err != nil {
		return destination, fmt.Errorf("failed to create file TESTER.txt %v", err)
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		//return fmt.Errorf("failed to download file, %T -- %v", err, err)
		return destination, fmt.Errorf("failed to download file, %T -- %v", err, err)
	}


	fmt.Printf("file WRITTEN TO TESTER.txt -- downloaded from %s :: %d bytes\n", destination, n)

	return destination, err
}





