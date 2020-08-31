package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/service/edit"
	"net/http"
)

func EditProfile(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	email := r.Header.Get("X-Consumer-Username")

	e := edit.Edit{}
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.Do(i, email)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}