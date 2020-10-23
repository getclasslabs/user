package handler

import (
	"encoding/json"
	"github.com/getclasslabs/go-tools/pkg/request"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
	"github.com/getclasslabs/user/internal/service/search"
	"net/http"
)

func SearchTeacher(w http.ResponseWriter, r *http.Request) {
	var retStatus int
	var retMessage string

	i := r.Context().Value(request.ContextKey).(*tracer.Infos)
	i.TraceIt(spanName)
	defer i.Span.Finish()

	name := r.URL.Query().Get("name")
	page := r.URL.Query().Get("page")

	if name == "" {
		i.Span.SetTag("no name", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if page == "" {
		page = "1"
	}

	result, err := search.Teacher(i, name, page)
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

	if len(result["results"].([]map[string]interface{})) == 0{
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ret, _ := json.Marshal(result)
	_, _ = w.Write(ret)
}
