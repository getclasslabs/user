package handler

import (
	"encoding/json"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/service/user"
	"github.com/getclasslabs/user/tools"
	"net/http"
)

//CreateUser Route for user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var retStatus int
	var retMessage string

	i := r.Context().Value(tools.ContextKey).(*tools.Infos)

	i.Span = tools.TraceIt(i, spanName)
	defer i.Span.Finish()
	createUserService := user.CreateUserService{}
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
