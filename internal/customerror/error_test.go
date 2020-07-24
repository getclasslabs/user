package customerror

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestError_Error(t *testing.T) {
	type fields struct {
		Origin     string
		Complement string
		From       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"[customerror]:Error error from a specific struct",
			fields{
				"error for testing",
				"a random error",
				"customerror.randomStruct",
			},
			"[error] a random error | origin: error for testing | from: customerror.randomStruct | extra: null",
		},
		{
			"[customerror]:Error error from a specific func",
			fields{
				"error for testing func",
				"an error",
				"func(string)",
			},
			"[error] an error | origin: error for testing func | from: func(string) | extra: null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Error{
				origin:     tt.fields.Origin,
				complement: tt.fields.Complement,
				from:       tt.fields.From,
			}
			if got := c.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testFunc(e string) {
}

type randomStruct struct{}

func TestNewError(t *testing.T) {
	r := randomStruct{}
	type args struct {
		from       interface{}
		complement string
		err        error
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			"[customerror]:New a function from a specific struct",
			args{
				r,
				"a random error",
				errors.New("error for testing"),
			},
			errors.New("[error] a random error | origin: error for testing | from: customerror.randomStruct | extra: null"),
		},
		{
			"[customerror]:New from a specific func",
			args{
				testFunc,
				"a random error",
				errors.New("error for testing"),
			},
			errors.New("[error] a random error | origin: error for testing | from: func(string) | extra: null"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewError(tt.args.from, tt.args.complement, tt.args.err)
			assert.Error(t, err)
			assert.Equal(t, tt.expectedErr.Error(), err.Error())
		})
	}
}
