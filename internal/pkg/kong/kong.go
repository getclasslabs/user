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
	ErrorCreatingJwt = "creating jwt"
)
const contentType = "application/x-www-form-urlencoded"

type Kong struct {
	httpClient pkg.HttpClientInterface
}

type Claims struct {
	Iss string `json:"iss"`
	jwt.StandardClaims
}

type jwtResponse struct {
	ConsumerId string `json:"consumer_id"`
	Key        string `json:"key"`
	Secret     string `json:"secret"`
}

func NewKong() *Kong {
	k := &Kong{
		&http.Client{},
	}
	return k
}

//CreateCustomer Creates a customer on kong POST to `http://konghost:8001/consumer`
func (k *Kong) CreateCustomer(email string) error {
	cfg := config.Config.Kong
	urlRequest := cfg.Host + cfg.ConsumerRequest
	data := url.Values{}
	data.Set("username", email)

	r, err := http.NewRequest(http.MethodPost, urlRequest, strings.NewReader(data.Encode()))

	if err != nil {
		return customerror.NewRequestError(k, err, urlRequest, 0)
	}

	r.Header.Set("Content-Type", contentType)
	resp, err := k.httpClient.Do(r)
	if err != nil {
		return customerror.NewRequestError(k, err, urlRequest, resp.StatusCode)
	}

	if resp.StatusCode != 201 {
		return customerror.NewRequestError(k, err, urlRequest, resp.StatusCode)
	}

	return nil
}

//CreateCredentials Creates a JWT keys to a user POST to `http://konghost:8001/consumers/%s/jwt`
func (k *Kong) CreateCredentials(email string) (string, error) {
	cfg := config.Config.Kong
	urlReq := cfg.Host + fmt.Sprintf(cfg.JwtRequest, email)

	r, err := http.NewRequest(http.MethodPost, urlReq, nil)
	if err != nil {
		return "", customerror.NewRequestError(k, err, urlReq, 0)
	}

	r.Header.Set("Content-Type", contentType)
	resp, err := k.httpClient.Do(r)
	if err != nil {
		return "", customerror.NewRequestError(k, err, urlReq, resp.StatusCode)
	}

	if resp.StatusCode != 201 {
		return "", customerror.NewRequestError(k, nil, urlReq, resp.StatusCode)
	}

	defer resp.Body.Close()
	jwtR := jwtResponse{}

	err = json.NewDecoder(resp.Body).Decode(&jwtR)
	if err != nil {
		return "", customerror.NewParsingError(k, err)
	}

	jwtCode, err := k.createJWT(&jwtR)

	return jwtCode, nil
}

func (k *Kong) createJWT(response *jwtResponse) (string, error) {
	claims := &Claims{
		Iss: response.Key,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(response.Secret))

	if err != nil {
		return "", customerror.NewError(k, ErrorCreatingJwt, err)
	}

	return tokenString, nil
}

var Service = NewKong()
