package utils

type Response struct {
	Status  status  `json:"status"`
	Message *string `json:"message,omitempty"`
	Data    any     `json:"data,omitempty"`
}

type status string

var (
	successStatus status = "Success"
	failedStatus  status = "Failed"
)

func NewError(err error) *Response {
	errMessage := err.Error()
	return &Response{
		Status:  failedStatus,
		Message: &errMessage,
	}
}

func NewSuccessResp(data any, messages ...string) *Response {
	var successMessage *string
	if len(messages) > 0 {
		successMessage = &messages[0]
	}
	return &Response{
		Status:  successStatus,
		Message: successMessage,
		Data:    data,
	}
}
