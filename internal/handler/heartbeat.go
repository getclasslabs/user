package handler

import (
	"encoding/json"
	"github.com/getclasslabs/user/tools"
	"net/http"
)

//Heartbeat only to check the health of the API
func Heartbeat(w http.ResponseWriter, r *http.Request) {
	i := r.Context().Value(tools.ContextKey).(*tools.Infos)

	i.Span = tools.TraceIt(i, spanName)
	defer i.Span.Finish()

	ret, _ := json.Marshal(map[string]string{
		"msg": "hello é o caralho",
	})
	_, _ = w.Write(ret)

}
