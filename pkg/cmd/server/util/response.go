package util

func JSONResponse(success bool, message string, data any) map[string]any {
	return map[string]any{
		"success": success,
		"message": message,
		"data":    data,
	}
}
