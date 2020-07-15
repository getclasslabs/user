package customerror

import (
	"fmt"
	"reflect"
)

type CustomError struct {
	Origin string
	Complement string
	From string
}

func NewError(from interface{}, complement string, err error) error {
	var f string
	f = reflect.TypeOf(from).String()
	return &CustomError{
		From:       f,
		Complement: complement,
		Origin:     err.Error(),
	}
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("[error] %s | origin: %s | from: %s", c.Complement, c.Origin, c.From)
}