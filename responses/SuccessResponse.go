package responses

import "encoding/json"

type SuccessResponse struct {
	Response string
}

func NewSuccessResponse(res string) string {
	resp := &SuccessResponse{res}
	data, _ := json.MarshalIndent(resp, "", "    ")
	return string(data)
}
