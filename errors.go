package nutshttp

type APIError struct {
	Code    int
	Message string
}

var (
	ErrNotFound            = APIError{404, "Not Found"}
	ErrInternalServerError = APIError{500, "Internal Server Error"}
)
