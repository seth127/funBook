package main

import (
	"fmt"
	"github.com/seth127/funBook/fbutils"
	"github.com/seth127/funBook/funbook"
	"strings"
	"time"

	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// This all needs to go in an aws module

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToAllToolsS3(s *session.Session, filePath string, destPath string) error {

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
		Bucket:               aws.String(fbutils.S3_BUCKET),
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

//////////

func main() {
	pick, pickPath := funbook.GetPickPath()

	fmt.Printf("\n ------ %d -------\n", pick)

	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{Region: aws.String(fbutils.S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// build destination path
	pathSlice := strings.Split(pickPath, "/")
	timeStamp := time.Now().Format("20060102-150405")

	destFile := pathSlice[len(pathSlice)-1] + "-" + timeStamp

	destPath := filepath.Join(fbutils.S3_MD_TEST_KEY, destFile)

	// Upload
	err = AddFileToAllToolsS3(s, pickPath, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully wrote to s3://%s/%s\n", fbutils.S3_BUCKET, destPath)

}