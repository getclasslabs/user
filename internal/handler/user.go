package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/service/userService"
	"github.com/gorilla/mux"
	"net/http"
)

//CreateUser Route for user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var retStatus int
	var retMessage string

	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	createUserService := userService.CreateUserService{}
	err := json.NewDecoder(r.Body).Decode(&createUserService)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwt, err := createUserService.Do(i)
	if err != nil {
		switch err.(type) {
		case customerror.CustomError:
			retStatus = customerror.HandleStatus(err.(customerror.CustomError))
			retMessage = err.(customerror.CustomError).GetMessage()
		default:
			retStatus = http.StatusInternalServerError
			retMessage = customerror.GenericErrorMessage
		}
		i.Span.SetTag("creating", http.StatusInternalServerError)
		w.WriteHeader(retStatus)
		_, _ = w.Write([]byte(retMessage))
		return
	}

	ret, _ := json.Marshal(map[string]string{
		"jwt": jwt,
	})
	_, _ = w.Write(ret)

}

// CreateProfile Creates user profile information
func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var retStatus int
	var retMessage string

	email := r.Header.Get("X-Consumer-Username")

	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	profileService := userService.Profile{}
	err := json.NewDecoder(r.Body).Decode(&profileService)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = profileService.Do(i, email)
	if err != nil {
		switch err.(type) {
		case customerror.CustomError:
			retStatus = customerror.HandleStatus(err.(customerror.CustomError))
			retMessage = err.(customerror.CustomError).GetMessage()
		default:
			retStatus = http.StatusInternalServerError
			retMessage = customerror.GenericErrorMessage
		}
		i.Span.SetTag("creating", http.StatusInternalServerError)
		w.WriteHeader(retStatus)
		_, _ = w.Write([]byte(retMessage))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var retStatus int
	var retMessage string

	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	nickname := mux.Vars(r)["nickname"]

	if len(nickname) == 0 {
		msg, _ := json.Marshal(map[string]string{"msg": "No nickname provided"})
		_, _ = w.Write(msg)
		w.WriteHeader(http.StatusBadRequest)
	}

	u, err := userService.Get(i, nickname)
	if err != nil {
		switch err.(type) {
		case customerror.CustomError:
			retStatus = customerror.HandleStatus(err.(customerror.CustomError))
			retMessage = err.(customerror.CustomError).GetMessage()
		default:
			retStatus = http.StatusInternalServerError
			retMessage = customerror.GenericErrorMessage
		}
		i.Span.SetTag("creating", http.StatusInternalServerError)
		w.WriteHeader(retStatus)
		_, _ = w.Write([]byte(retMessage))
		return
	}

	ret, _ := json.Marshal(u)
	_, _ = w.Write(ret)

}
