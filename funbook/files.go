package funbook

import (
	"os"
	"io/ioutil"
	"github.com/seth127/funBook/fbutils"
	"path/filepath"
)

func WritePickLocal(s string, n int, outDir string) {

	filename := filepath.Join(outDir, fbutils.PadNumberWithZero(uint32(n)))

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if _, ok := err.(*os.PathError); ok {
		f, err = os.Create(filename)
	}

	fbutils.PanicNil(err)

	defer f.Close()

	_, err = f.WriteString(s + "\n")
	fbutils.PanicNil(err)
}

func ReadPickLocal(path string, n int) (string, string, error) {

	filename := filepath.Join(path, fbutils.PadNumberWithZero(uint32(n)))

	dat, err := ioutil.ReadFile(filename)

	return string(dat), filename, err
}
