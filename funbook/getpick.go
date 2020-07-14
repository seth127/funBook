package funbook

import (
	"os"

	"github.com/seth127/funBook/fbutils"
)


func GetPickString() (int, string) {
	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	pickString, _, pickErr := ReadParagraph(fbutils.OutDir, pick)
	fbutils.CheckNonMissing(pickErr)
	_, missingErr := pickErr.(*os.PathError)

	for missingErr {
		pick = fbutils.PickRand(fbutils.MaxParagraphs)
		pickString, _, pickErr = ReadParagraph(fbutils.OutDir, pick)
		fbutils.CheckNonMissing(pickErr)
		_, missingErr = pickErr.(*os.PathError)
	}

	return pick, pickString
}

func GetPickPath() (int, string) {
	pick := fbutils.PickRand(fbutils.MaxParagraphs)
	_, pickPath, pickErr := ReadParagraph(fbutils.OutDir, pick)
	fbutils.CheckNonMissing(pickErr)
	_, missingErr := pickErr.(*os.PathError)

	for missingErr {
		pick = fbutils.PickRand(fbutils.MaxParagraphs)
		_, pickPath, pickErr = ReadParagraph(fbutils.OutDir, pick)
		fbutils.CheckNonMissing(pickErr)
		_, missingErr = pickErr.(*os.PathError)
	}

	return pick, pickPath
}