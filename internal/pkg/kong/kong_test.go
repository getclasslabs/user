package kong

import (
	"bytes"
	"github.com/getclasslabs/user/internal/pkg"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type mockedHttp struct {
	returnCode int
	err        error
	resp       string
}

func (m *mockedHttp) Do(_ *http.Request) (*http.Response, error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(m.resp)))
	return &http.Response{
		StatusCode: m.returnCode,
		Body:       r,
	}, m.err
}

func TestKong_CreateCustomer(t *testing.T) {
	type fields struct {
		httpClient pkg.HttpClientInterface
	}
	tests := []struct {
		name    string
		fields  fields
		email   string
		wantErr bool
	}{
		{
			"[Kong] 201 from kong",
			fields{
				httpClient: &mockedHttp{
					201,
					nil,
					"",
				},
			},
			"and_lrg@hotmail.com",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kong{
				httpClient: tt.fields.httpClient,
			}
			if err := k.CreateCustomer(tt.email); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKong_CreateCredentials(t *testing.T) {
	type fields struct {
		httpClient pkg.HttpClientInterface
		cType      string
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kong{
				httpClient: tt.fields.httpClient,
			}
			got, err := k.CreateCredentials(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateCredentials() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKong_createJWT(t *testing.T) {
	type fields struct {
		httpClient pkg.HttpClientInterface
		cType      string
	}
	type args struct {
		response *jwtResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Kong{
				httpClient: tt.fields.httpClient,
			}
			got, err := k.createJWT(tt.args.response)
			if (err != nil) != tt.wantErr {
				t.Errorf("createJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKong(t *testing.T) {
	tests := []struct {
		name string
		want *Kong
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKong(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKong() = %v, want %v", got, tt.want)
			}
		})
	}
}
