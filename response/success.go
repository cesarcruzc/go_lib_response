package response

import (
	"encoding/json"
	"github.com/cesarcruzc/go_meta/meta"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

func OK(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusOK)
}

func Created(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusCreated)
}

func Accepted(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusAccepted)
}

func NonAuthoritativeInfo(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNoContent)
}

func ResetContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusResetContent)
}

func PartialContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusPartialContent)
}

func success(message string, data interface{}, meta *meta.Meta, status int) Response {
	return &SuccessResponse{
		Message: message,
		Status:  status,
		Data:    data,
		Meta:    meta,
	}
}

func (s *SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse) GetData() interface{} {
	return s.Data
}
