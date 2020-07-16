package tools

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

const ContextKey = "infos"

type controller func(http.ResponseWriter, *http.Request)

type Infos struct {
	Span   opentracing.Span
	Tracer opentracing.Tracer
}

func (i *Infos) LogError(err error){
	i.Span.SetTag("error", true)
	i.Span.SetTag("errorMsg", err.Error())
}

func PreRequest(h controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()
		tCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		span := tracer.StartSpan(fmt.Sprintf("%s %s", r.Method, r.URL.Path), ext.RPCServerOption(tCtx))
		defer span.Finish()
		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(span, r.URL.String())

		ctx := context.WithValue(r.Context(), ContextKey, &Infos{span, tracer})

		w.Header().Set("Content-Type", "application/json")
		h(w, r.WithContext(ctx))
	}
}
