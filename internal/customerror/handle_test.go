package customerror

import (
	"errors"
	"net/http"
	"testing"
)

func TestHandleStatus(t *testing.T) {
	type args struct {
		err CustomError
	}
	err := errors.New("testing")
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"[handleErrorStatusCode] request error 1",
			args{
				NewRequestError(TestHandleStatus, err, "testURL", 400).(CustomError),
			},
			http.StatusBadRequest,
		},
		{
			"[handleErrorStatusCode] request error 2",
			args{
				NewRequestError(TestHandleStatus, err, "testURL", 502).(CustomError),
			},
			http.StatusBadGateway,
		},
		{
			"[handleErrorStatusCode] parsing",
			args{
				NewParsingError(TestHandleStatus, err).(CustomError),
			},
			http.StatusBadRequest,
		},
		{
			"[handleErrorStatusCode] default",
			args{
				NewError("", "", err).(CustomError),
			},
			http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleStatus(tt.args.err); got != tt.want {
				t.Errorf("HandleStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
