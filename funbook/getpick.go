package funbook

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"path/filepath"

	"github.com/seth127/funBook/fbutils"

)


func GetPickLocal() (int, string, string) {
	MAX_TRIES := 100

	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	pickString, pickPath, pickErr := ReadPickLocal(fbutils.OutDir, pick)
	fbutils.CheckNonMissing(pickErr)
	_, missingErr := pickErr.(*os.PathError)

	for i := 1; missingErr; i++ {
		pick = fbutils.PickRand(fbutils.MaxParagraphs)
		pickString, pickPath, pickErr = ReadPickLocal(fbutils.OutDir, pick)
		fbutils.CheckNonMissing(pickErr)
		_, missingErr = pickErr.(*os.PathError)

		if i >= MAX_TRIES {
			panic(fmt.Sprintf("GetPickLocal() tried %d times and failed to find pick.", i))
		}
	}

	return pick, pickString, pickPath
}


func GetPickS3(s *session.Session) (int, string, error) {
	//MAX_TRIES := 100

	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	fmt.Printf("\n ------ %d -------\n", pick)
	pickPath := filepath.Join(fbutils.S3_OUT_KEY, fbutils.PadNumberWithZero(uint32(pick)))
	destination, pickErr := ReadFileFromS3(s, fbutils.S3_BUCKET, pickPath)
	//fbutils.CheckNonMissing(pickErr)
	//_, missingErr := pickErr.(*s3err.RequestFailure)
	fmt.Printf("NAW %T -- %v PAW\n", pickErr, pickErr)

	//for i := 1; missingErr; i++ {
	//	fmt.Printf("\n ------ %d -------\n", pick)
	//	pick = fbutils.PickRand(fbutils.MaxParagraphs)
	//	pickPath := filepath.Join(fbutils.S3_OUT_KEY, fbutils.PadNumberWithZero(uint32(pick)))
	//	//pickString, pickPath,
	//	pickErr = ReadFileFromS3(s, fbutils.OutDir, pickPath)
	//	fbutils.CheckNonMissing(pickErr)
	//	_, missingErr = pickErr.(*s3err.RequestFailure)
	//
	//	if i >= MAX_TRIES {
	//		panic(fmt.Sprintf("GetPickLocal() tried %d times and failed to find pick.", i))
	//	}
	//}

	//return pick, pickString, pickPath
	return pick, destination, pickErr
}