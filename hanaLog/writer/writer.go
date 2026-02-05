package writer

type Writer interface {
	Write(p []byte) error
}
