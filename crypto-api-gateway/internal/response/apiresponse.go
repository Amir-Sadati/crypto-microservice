package response

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiResponse[T any] struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Data       *T     `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
}

type ApiResponseNoData struct {
	StatusCode int    `json:"status_code"`
	Success    bool   `json:"success"`
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
}

func Success[T any](data T, msg string) ApiResponse[T] {
	return ApiResponse[T]{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    msg,
		Data:       &data,
	}
}

func Fail[T any](data T, statusCode int, err, msg string) ApiResponse[T] {
	return ApiResponse[T]{
		StatusCode: statusCode,
		Success:    false,
		Message:    msg,
		Error:      err,
		Data:       &data,
	}
}

// non-generic
func SuccessNoData(msg string) ApiResponseNoData {
	return ApiResponseNoData{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    msg,
	}
}

func FailNoData(status int, err, msg string) ApiResponseNoData {
	return ApiResponseNoData{
		StatusCode: status,
		Success:    false,
		Error:      err,
		Message:    msg,
	}
}

// ----------------------------
// gRPC Error Handler
// ----------------------------

func WriteFailFromGrpc(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	errMsg := "unexpected error"

	if s, ok := status.FromError(err); ok {
		statusCode, errMsg = mapGrpcToHttp(s.Code(), s.Message())
		if errMsg == "" {
			errMsg = s.Message()
		}
	}

	res := ApiResponse[any]{
		StatusCode: statusCode,
		Success:    false,
		Error:      errMsg,
	}

	writeJson(w, res.StatusCode, res)
}

// ----------------------------
// gRPC to HTTP Mapping
// ----------------------------
func mapGrpcToHttp(code codes.Code, msg string) (int, string) {
	switch code {
	case codes.NotFound:
		return http.StatusNotFound, msg
	case codes.InvalidArgument:
		return http.StatusBadRequest, msg
	case codes.AlreadyExists:
		return http.StatusConflict, msg
	case codes.Unauthenticated:
		return http.StatusUnauthorized, "Authentication required"
	case codes.PermissionDenied:
		return http.StatusForbidden, "Access denied"
	case codes.Internal:
		return http.StatusInternalServerError, "Internal server error"
	case codes.Unavailable:
		return http.StatusServiceUnavailable, "Service unavailable"
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout, "Request timeout"
	default:
		return http.StatusInternalServerError, "Unexpected error"
	}
}

// ----------------------------
// Write JSON Response
// ----------------------------

func WriteSuccessNoData(w http.ResponseWriter, msg string) {
	res := SuccessNoData(msg)
	writeJson(w, res.StatusCode, res)
}

func WriteSuccess[T any](w http.ResponseWriter, data T, msg string) {
	res := Success(data, msg)
	writeJson(w, res.StatusCode, res)
}

func WriteFailNoData(w http.ResponseWriter, statusCode int, err string, msg string) {
	res := FailNoData(statusCode, err, msg)
	writeJson(w, res.StatusCode, res)
}

func WriteFailWithData[T any](w http.ResponseWriter, data T, statusCode int, err string, msg string) {
	res := Fail(data, statusCode, err, msg)
	writeJson(w, res.StatusCode, res)
}

func writeJson(w http.ResponseWriter, status int, data any, headers ...http.Header) {
	out, err := json.Marshal(data)
	if err != nil {
		return
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(out)
}
