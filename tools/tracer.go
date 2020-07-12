package tools

import ot "github.com/opentracing/opentracing-go"

func TraceIt(i *Infos, name string) ot.Span {
	return i.Tracer.StartSpan(name, ot.ChildOf(i.Span.Context()))
}