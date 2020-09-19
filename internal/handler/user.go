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

func UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	email := r.Header.Get("X-Consumer-Username")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"msg": "The image sent is bigger than 10mb"}`))
	}

	file, _, err := r.FormFile("photo")
	if err != nil {
		i.Span.SetTag("getting form file", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = userService.UpdateImage(i, email, file)
	if err != nil {
		i.Span.SetTag("updating image", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	email := r.Header.Get("X-Consumer-Username")

	err := userService.ErasePhoto(i, email)
	if err != nil {
		i.Span.SetTag("getting form file", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

