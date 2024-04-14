package valueObject

type FilePath struct {
	Value string
}

func NewFilePath(path string) *FilePath {
	return &FilePath{
		Value: path,
	}
}
