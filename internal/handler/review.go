package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/service/review"
	"github.com/gorilla/mux"
	"net/http"
)

func Review(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	reviewService := review.Create{}
	err := json.NewDecoder(r.Body).Decode(&reviewService)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = reviewService.Create(i)
	if err != nil {
		i.Span.SetTag("read", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	teacherId := mux.Vars(r)["teacher"]

	if len(teacherId) == 0 {
		msg, _ := json.Marshal(map[string]string{"msg": "No teacher id provided"})
		_, _ = w.Write(msg)
		w.WriteHeader(http.StatusBadRequest)
	}

	u, err := review.Get(i, teacherId)
	if err != nil {
		i.Span.SetTag("creating", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(customerror.GenericErrorMessage))
		return
	}

	ret, _ := json.Marshal(u)
	_, _ = w.Write(ret)
}
