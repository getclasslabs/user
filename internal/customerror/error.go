package customerror

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type CustomError interface {
	Error() string
	GetErrType() string
	GetExtraInfo() map[string]interface{}
	GetMessage() string
}

type Error struct {
	origin     string
	complement string
	from       string
	errType    string
	message    string
	extraInfo  map[string]interface{}
}

func (c *Error) GetErrType() string {
	return c.errType
}

func (c *Error) GetExtraInfo() map[string]interface{} {
	return c.extraInfo
}

func (c *Error) Error() string {
	e, _ := json.Marshal(c.extraInfo)
	return fmt.Sprintf("[error] %s | origin: %s | from: %s | extra: %s", c.complement, c.origin, c.from, e)
}

func (c *Error) GetMessage() string {
	return c.message
}

func newError(from interface{}, complement string, err error, errType, message string, extra map[string]interface{}) error {
	var f string
	f = reflect.TypeOf(from).String()
	return &Error{
		from:       f,
		complement: complement,
		origin:     err.Error(),
		message:    message,
		errType:    errType,
		extraInfo:  extra,
	}
}
func NewError(from interface{}, complement string, err error) error {
	return newError(from, complement, err, Default, "A generic error occurred", nil)
}

func NewRequestError(from interface{}, err error, url string, statusCode int) error {
	return newError(
		from,
		"doing request",
		err,
		Request,
		"An internal request error happened",
		map[string]interface{}{
			"url":         url,
			"status_code": statusCode,
		})
}

func NewParsingError(from interface{}, err error) error {
	return newError(from, "parsing error", err, Parsing, "Some information is wrong", nil)
}
