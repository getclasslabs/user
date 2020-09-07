package domains

import "github.com/getclasslabs/go-tools/pkg/tracer"

type Domain struct {
	Tracer *tracer.Infos
	Email  string
}
