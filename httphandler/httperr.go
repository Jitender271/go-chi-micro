package httphandler

import (
    "errors"
    "fmt"
    "github.com/go-chi/render"
    "github.com/go-chi-micro/model"
    "net/http"
)

type HTTPErr struct {
    Err  error
    Code int
}

func (e HTTPErr) Error() string {
    return fmt.Sprintf("%s", e.Err)
}

// New returns new http error from error object and a code
func New(err error, code int) *HTTPErr {
    return &HTTPErr{
        Err:  err,
        Code: code,
    }
}

// Error returns an HTTPError
func Error(err interface{}) *HTTPErr {
    switch err.(type) {
    case *HTTPErr:
        return err.(*HTTPErr)
    case error:
        return &HTTPErr{
            Err:  err.(error),
            Code: http.StatusInternalServerError,
        }
    default:
        return &HTTPErr{
            Err:  errors.New("Unknown error"),
            Code: http.StatusInternalServerError,
        }
    }
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, e.HTTPStatusCode)
    return nil
}

//Render for All Responses
func (rd *Response) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

// Response is a wrapper response structure
type Response struct {
    Status interface{} `json:"status"`
    Data   interface{} `json:"data, omitempty"`
}

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
    HTTPStatusCode int                `json:"-"` // http response status code
    Status         model.ResponseMeta `json:"status"`
    AppCode        int64              `json:"code,omitempty"` // application-specific error code
}

func WrapHandlerFunc(route string, name string, handler http.HandlerFunc) (string, http.HandlerFunc) {
    return route, handler
}

// NewSuccessResponse returns a new success response
func NewSuccessResponse(status int, data interface{}) *Response {
    return &Response{
        Status: &model.ResponseMeta{
            AppStatusCode: status,
            Message:       "SUCCESS",
        },
        Data: data,
    }
}

func ErrInvalidRequest(err error, message string) render.Renderer {
    return &ErrResponse{
        HTTPStatusCode: http.StatusOK,
        Status: model.ResponseMeta{
            AppStatusCode: http.StatusBadRequest,
            Message:       "ERROR",
            ErrorMessage:  "Invalid Request",
            ErrorDetail:   message,
            DevMessage:    err.Error(),
        },
    }
}

// ErrNotFound returns a 404
var ErrNotFound = &ErrResponse{
    HTTPStatusCode: http.StatusOK,
    Status: model.ResponseMeta{
        AppStatusCode: http.StatusNotFound,
        Message:       "ERROR",
        ErrorDetail:   "Resource not found",
        ErrorMessage:  "The endpoint you were seeking burned down",
    },
}
