package customerror

import (
	"fmt"
	"strconv"
)

type Request struct {
	Url        string
	StatusCode int
}

func NewRequestError(url string, statuscode int) error{
	return &Request{
		Url:        url,
		StatusCode: statuscode,
	}
}

func (r *Request) Error() string {
	return fmt.Sprintf("[error] request to %s returned %s", r.Url, strconv.Itoa(r.StatusCode))
}

