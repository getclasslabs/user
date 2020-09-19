package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/service/login"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	loginService := login.Login{}
	err := json.NewDecoder(r.Body).Decode(&loginService)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := loginService.Do(i)

	ret, _ := json.Marshal(result)
	_, _ = w.Write(ret)

}
