package funbook

import (
	"os"
	"io/ioutil"
	"github.com/seth127/funBook/fbutils"
)

func WriteParagraph(s string, n int, outDir string) {

	filename := outDir + fbutils.PadNumberWithZero(uint32(n))

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if _, ok := err.(*os.PathError); ok {
		f, err = os.Create(filename)
	}

	fbutils.CheckPanic(err)

	defer f.Close()

	_, err = f.WriteString(s + "\n")
	fbutils.CheckPanic(err)
}

func ReadParagraph(path string, n int) (string, error) {

	filename := path + fbutils.PadNumberWithZero(uint32(n))

	dat, err := ioutil.ReadFile(filename)

	return string(dat), err
}