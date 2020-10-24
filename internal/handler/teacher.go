package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/service/teacher"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SearchTeacherById(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		i.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t, err := teacher.GetTeacherById(i, id)

	ret, _ := json.Marshal(t)
	_, _ = w.Write(ret)
}
