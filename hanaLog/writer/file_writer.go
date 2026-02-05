package writer

import "os"

type FileWriter struct {
	file *os.File
}

func (w *FileWriter) Write(p []byte) error {
	_, err := w.file.Write(p)
	return err
}
