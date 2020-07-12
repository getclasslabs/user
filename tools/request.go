package tools

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

type controller func(http.ResponseWriter, *http.Request)

type Infos struct{
	Span opentracing.Span
	Tracer opentracing.Tracer
}

func PreRequest(h controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()
		tCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		span := tracer.StartSpan(fmt.Sprintf("%s %s", r.Method, r.URL.Path), ext.RPCServerOption(tCtx))
		defer span.Finish()
		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(span, r.URL.String())

		ctx := context.WithValue(r.Context(), "infos", &Infos{span, tracer})

		w.Header().Set("Content-Type", "application/json")
		h(w, r.WithContext(ctx))
	}
}