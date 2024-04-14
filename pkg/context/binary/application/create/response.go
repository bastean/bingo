package create

type Response struct {
	FilePath string
}

func NewResponse(filePath string) *Response {
	return &Response{
		FilePath: filePath,
	}
}
