package generators

import (
	"io/fs"
	"io/ioutil"
)

type Writer struct {}

type WriterInterface interface {
	WriteFile(filename string, data []byte, perm fs.FileMode) error
}

func (w *Writer) WriteFile(filename string, data []byte, perm fs.FileMode) error{
	return ioutil.WriteFile(filename, data, perm)
}