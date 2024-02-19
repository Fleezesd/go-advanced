package IO

import (
	"io"
	"os"
)

type MultiReaderAndWriter interface {
	Copy(dst Writer, src Reader) (written int64, err error)
	MultiReader(readers ...Reader) Reader
	MultiWriter(writers ...Writer) Writer
}

// MultiReaderToMultiWriter 将MultiReader, MultiWriter, Copy进行结合
func MultiReaderToMultiWriter() {
	var r1, r2, r3 *os.File
	var w1, w2, w3 *os.File
	_, _ = io.Copy(io.MultiWriter(w1, w2, w3), io.MultiReader(r1, r2, r3))
}
