package funbook

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seth127/funBook/fbutils"
	"os"
	"path/filepath"
)

// Picks a random paragraph from S3 and returns
// the number of the pick (int),
// the byte slice of the text retrieved from S3 ([]byte),
// and the S3 path that the slice came from (string).
// Note that it will panic on an unexpected error through PanicNonS3Missing().
// We may want to refactor this to return errors later, but given the missingErr pattern
// it was easier to do this.
func GetPickS3(s *session.Session) (int, []byte, string) {

	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	//fmt.Printf("\n ------ %d -------\n", pick)
	pickPath := filepath.Join(fbutils.S3_OUT_KEY, fbutils.PadNumberWithZero(uint32(pick)))
	pickBytes, pickErr := ReadFileFromS3(s, fbutils.S3_BUCKET, pickPath)
	fbutils.PanicNonS3Missing(pickErr)
	missingErr := fbutils.CheckS3MissingError(pickErr)

	for i := 1; missingErr; i++ {
		//fmt.Printf("\n ------ %d ------- REDO\n", pick)
		pick = fbutils.PickRand(fbutils.MaxParagraphs)
		pickPath = filepath.Join(fbutils.S3_OUT_KEY, fbutils.PadNumberWithZero(uint32(pick)))
		pickBytes, pickErr = ReadFileFromS3(s, fbutils.S3_BUCKET, pickPath)
		fbutils.PanicNonS3Missing(pickErr)
		missingErr = fbutils.CheckS3MissingError(pickErr)

		if i >= fbutils.MAX_TRIES {
			panic(fmt.Sprintf("GetPickS3() tried %d times and failed to find pick.", i))
		}
	}

	return pick, pickBytes, fmt.Sprintf("s3://%s/%s", fbutils.S3_BUCKET, pickPath)
}

// Picks a random paragraph from fbutils.OutDir and returns
// the number of the pick (int),
// the the text retrieved from the picked file (string),
// and the file path that the text came from (string).
// Note that it will panic on an unexpected error through PanicNonPathError().
// We may want to refactor this to return errors later, but given the missingErr pattern
// it was easier to do this.
func GetPickLocal() (int, string, string) {

	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	pickString, pickPath, pickErr := ReadPickLocal(fbutils.OutDir, pick)
	fbutils.PanicNonPathError(pickErr)
	_, missingErr := pickErr.(*os.PathError)

	for i := 1; missingErr; i++ {
		pick = fbutils.PickRand(fbutils.MaxParagraphs)
		pickString, pickPath, pickErr = ReadPickLocal(fbutils.OutDir, pick)
		fbutils.PanicNonPathError(pickErr)
		_, missingErr = pickErr.(*os.PathError)

		if i >= fbutils.MAX_TRIES {
			panic(fmt.Sprintf("GetPickLocal() tried %d times and failed to find pick.", i))
		}
	}

	return pick, pickString, pickPath
}

