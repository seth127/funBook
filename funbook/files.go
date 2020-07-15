package funbook

import (
	"os"
	"io/ioutil"
	"github.com/seth127/funBook/fbutils"
	"path/filepath"
)

// These might all get deprecated. They were for initial testing,
// although they are still being used by getbook.go
// until it gets refactored to write to s3.

// Write string to file in outDir with filename
// of n zero-padded to 5 digits.
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

// Read a file from directory passed to path
// and filename of n zero-padded to 5 digits.
func ReadPickLocal(path string, n int) (string, string, error) {

	filename := filepath.Join(path, fbutils.PadNumberWithZero(uint32(n)))

	dat, err := ioutil.ReadFile(filename)

	return string(dat), filename, err
}
