package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status int         `json:"status"`
	Header Header      `json:"header"`
	Data   interface{} `json:"data,omitempty"`
}

type Header struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

//NewRestResponse is to create new rest response
func NewRestResponse(startTime time.Time) Response {
	return Response{}
}

//WriteResponse to write rest regular response
func (res *Response) WriteResponse(w http.ResponseWriter, msg interface{}) {
	res.Data = msg
	encoded, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

//WriteError to write rest error response
func (res *Response) WriteError(w http.ResponseWriter, code int, msg interface{}) {
	res.Header.Message = msg
	encoded, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(encoded)
}
