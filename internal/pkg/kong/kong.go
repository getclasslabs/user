package kong

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/getclasslabs/user/internal/config"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/pkg"
	"net/http"
	"net/url"
	"strings"
)

const (
	ErrorCreatingBodyConsumer = "creating body consumer"
	ErrorCreatingRequestConsumer = "creating request consumer"
	ErrorDoingRequestConsumer = "doing request consumer"
	ErrorCreatingRequestCredentials = "creating request credentials"
	ErrorDoingRequestCredentials = "doing request credentials"
	ErrorDecodingResponse = "request credentials couldnt decode response"
	ErrorCreatingJwt = "creating jwt"
)

type Kong struct {
	httpClient pkg.HttpClientInterface
	cType string
}

type Claims struct {
	Iss string `json:"iss"`
	jwt.StandardClaims
}

type jwtResponse struct {
	ConsumerId string `json:"consumer_id"`
	Key string `json:"key"`
	Secret string `json:"key"`
}

func NewKong() *Kong {
	k := &Kong{
		&http.Client{},
		"application/x-www-form-urlencoded",
	}
	return k
}

func (k *Kong) CreateCustomer(email string) error {
	cfg := config.Config.Kong
	urlRequest := cfg.Host + cfg.ConsumerRequest
	data := url.Values{}
	data.Set("username", email)

	r, err := http.NewRequest(http.MethodPost, urlRequest, strings.NewReader(data.Encode()))

	if err != nil{
		return customerror.NewError(k, ErrorCreatingRequestConsumer, err)
	}

	r.Header.Set("Content-Type", k.cType)
	resp, err := k.httpClient.Do(r)
	if err != nil{
		return customerror.NewError(k, ErrorDoingRequestConsumer, err)
	}

	if resp.StatusCode != 201 {
		return customerror.NewRequestError(urlRequest, resp.StatusCode)
	}

	return nil
}


func (k *Kong) CreateCredentials(email string) (string, error) {
	cfg := config.Config.Kong
	urlReq := cfg.Host + fmt.Sprintf(cfg.JwtRequest, email)

	r, err := http.NewRequest(http.MethodPost, urlReq, nil)
	if err != nil{
		return "", customerror.NewError(k, ErrorCreatingRequestCredentials, err)
	}

	r.Header.Set("Content-Type", k.cType)
	resp, err := k.httpClient.Do(r)
	if err != nil {
		return "", customerror.NewError(k, ErrorDoingRequestCredentials, err)
	}

	if resp.StatusCode != 201 {
		return "", customerror.NewRequestError(urlReq, resp.StatusCode)
	}

	defer resp.Body.Close()
	jwtR := jwtResponse{}

	err = json.NewDecoder(resp.Body).Decode(&jwtR)
	if err != nil {
		return "", customerror.NewError(k, ErrorDecodingResponse, err)
	}

	jwtCode, err := k.createJWT(&jwtR)


	return jwtCode, nil
}

func (k *Kong) createJWT(response *jwtResponse) (string, error) {
	claims := &Claims{
		Iss: response.Secret,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(response.Key))

	if err != nil {
		return "", customerror.NewError(k, ErrorCreatingJwt, err)
	}

	return tokenString, nil
}