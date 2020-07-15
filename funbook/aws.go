package funbook

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/seth127/funBook/fbutils"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// This is a placeholder to write the byte array of the picked paragraph to a text file in S3,
// but it will be replaced by something to write the paragraph into an email instead.
func AddPickToS3(s *session.Session, pickBytes []byte, pickPath string) {
	// build destination path
	pathSlice := strings.Split(pickPath, "/")
	timeStamp := time.Now().Format("20060102-150405")
	destFile :=  fmt.Sprintf("%s-%s", timeStamp, pathSlice[len(pathSlice)-1])
	destPath := filepath.Join(fbutils.S3_MD_TEST_KEY, destFile)

	// Upload
	err := WriteBufferToS3(s, fbutils.S3_BUCKET, pickBytes, destPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully wrote to s3://%s/%s\n", fbutils.S3_BUCKET, destPath)
}

// WriteBufferToS3 uploads a buffer (byte slice) to a file in S3.
// The `destPath` argument is often called a `key` in S3 terminology.
// Destination file will be at s3://[bucket]/[destPath]
// This function requires a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func WriteBufferToS3(s *session.Session, bucket string, buffer []byte, destPath string) error {

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(destPath),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(len(buffer))),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	return err
}

// Reads a file from S3 into a buffer.
// This function requires a pre-built aws session.
func ReadFileFromS3(s *session.Session, bucket string, key string) ([]byte, error) {

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(s)

	// Write the contents of S3 Object to a buffer
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		//return fmt.Errorf("failed to download file, %T -- %v", err, err)
		return make([]byte, 0), fmt.Errorf("failed to download file, %T -- %v", err, err)
	}

	return buf.Bytes(), err
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
// NOTE: This was the original boilerplate I copied from https://golangcode.com/uploading-a-file-to-s3/
// It is _not_ currently used in the package because it was adapted into WriteBufferToS3() for our purposes.
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
