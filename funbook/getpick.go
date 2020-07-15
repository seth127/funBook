package funbook

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seth127/funBook/fbutils"
	"os"
	"path/filepath"
)


func GetPickLocal() (int, string, string) {
	MAX_TRIES := 100

	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	pickString, pickPath, pickErr := ReadPickLocal(fbutils.OutDir, pick)
	fbutils.PanicNonPathError(pickErr)
	_, missingErr := pickErr.(*os.PathError)

	for i := 1; missingErr; i++ {
		pick = fbutils.PickRand(fbutils.MaxParagraphs)
		pickString, pickPath, pickErr = ReadPickLocal(fbutils.OutDir, pick)
		fbutils.PanicNonPathError(pickErr)
		_, missingErr = pickErr.(*os.PathError)

		if i >= MAX_TRIES {
			panic(fmt.Sprintf("GetPickLocal() tried %d times and failed to find pick.", i))
		}
	}

	return pick, pickString, pickPath
}


func GetPickS3(s *session.Session) (int, []byte, string) {
	MAX_TRIES := 100

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

		if i >= MAX_TRIES {
			panic(fmt.Sprintf("GetPickS3() tried %d times and failed to find pick.", i))
		}
	}

	return pick, pickBytes, fmt.Sprintf("s3://%s/%s", fbutils.S3_BUCKET, pickPath)
}