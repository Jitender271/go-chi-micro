package model

type RecordSchema struct {
    ID        string `json:"id"`
    Content   string `json:"content"`
    CreatedAt string `json:"created_at"`
    Title     string `json:"title"`
    UpdatedAt string `json:"updated_at"`
}

type SelectService struct {
    Team  string `json:"team"`
    Email string `json:"email"`
}

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

type Service struct {
    Name         string `json:"name"`
    InstanceType string `json:"instance_type,omitempty"`
    Count        int    `json:"count,omitempty"`
    Build        string `json:"build,omitempty"`
    AccessKey    string `json:"access_key,omitempty"`
    SecretKey    string `json:"secret_key,omitempty"`
    RunId        string `json:"runId,omitempty"`
    Pod          string `json:"pod,omitempty"`
}

type Blog struct{
    ID  int `json:"id"`
    Blogname string `json:"blogname"`
    BlogDetails string `json:"blog_details"`
    Email string `json:"email"`
}