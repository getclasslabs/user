package handler

import (
	"encoding/json"
	"github.com/getclasslabs/user/internal/service/user"
	"github.com/getclasslabs/user/tools"
	"net/http"
)

//CreateUser Route for user creation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(tools.ContextKey).(*tools.Infos)

	i.Span = tools.TraceIt(i, spanName)
	defer i.Span.Finish()
	ur := user.Request{}
	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil{
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwt, err := user.CreateUser(i, &ur)
	if err != nil{
		//TODO Handle errors to fit REST
		i.Span.SetTag("creating", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ret, _ := json.Marshal(map[string]string{
		"jwt": jwt,
	})
	_, _ = w.Write(ret)

}
