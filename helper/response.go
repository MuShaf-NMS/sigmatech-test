package helper

// Helper to generate error response
func ErrorResponseBuilder(errs interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status": "Failed",
		"err":    errs,
	}
	return res

}

// Helper to generate respons
func ResponseBuilder(data interface{}) map[string]interface{} {
	res := map[string]interface{}{
		"status": "OK",
		"data":   data,
	}
	return res
}
