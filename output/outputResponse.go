package output

import "encoding/json"

type RestSucess struct { // Success Message Struct
	Message    string      `json:"message"`
	Value      interface{} `json:"data"`
	StatusCode int         `json:"status"`
	Type       string      `json:"type"`
}

/*
Success Output Function
Arguments : message string -
print the Success Message and value interface -
print the Multiple Values
*/
func SuccessRequest(message string, value interface{}, status int) *RestSucess {
	return &RestSucess{
		Message:    message,
		Value:      value,
		StatusCode: status,
		Type:       "Success",
	}
}

type Resterr struct { // Error Message Struct
	Message    string      `json:"message"`
	Value      interface{} `json:"data"`
	StatusCode int         `json:"status"`
	Type       string      `json:"type"`
}

/*
Error Output Function
Arguments : message string -
print the Error Message
*/
func ErrorRequest(message, errorType string, status int) *Resterr {
	return &Resterr{
		Message:    message,
		Value:      json.RawMessage(`{}`),
		StatusCode: status,
		Type:       errorType,
	}
}
