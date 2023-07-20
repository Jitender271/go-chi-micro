package model


type ResponseMeta struct {
    AppStatusCode int    `json:"code"`
    Message       string `json:"statusType,omitempty"`
    ErrorDetail   string `json:"errorDetail,omitempty"`
    ErrorMessage  string `json:"errorMessage,omitempty"`
    DevMessage    string `json:"devErrorMessage,omitempty"`
}

type ErrResponse struct {
    HTTPStatusCode int          `json:"-"` // http response status code
    Status         ResponseMeta `json:"status"`
    AppCode        int64        `json:"code,omitempty"` // application-specific error code
}

type Blogs struct{
    ID  int `json:"id"`
    BlogName string `json:"blog_name"`
    BlogDetails string `json:"blog_details"`
    Email string `json:"email"`
}