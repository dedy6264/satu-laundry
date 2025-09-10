package entities

type (
	APIResponse struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Error   string      `json:"error,omitempty"`
	}
	ResponseGlobal struct {
		ResponseCode    string      `json:"responseCode"`
		ResponseMessage string      `json:"responseMessage"`
		Result          interface{} `json:"result"`
	}
)
