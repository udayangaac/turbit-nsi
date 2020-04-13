package file_manager

type FileManager interface {
	Read(path string, i interface{}) (readErr error)
	Write(path string, i interface{}) (err error)
}
