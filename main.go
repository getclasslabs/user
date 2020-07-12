package main

import (
	"github.com/getclasslabs/user/internal"
	"github.com/getclasslabs/user/internal/config"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConf "github.com/uber/jaeger-client-go/config"
	jaegerLog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

func main() {

	cfg := jaegerConf.Configuration{
		ServiceName: "go-chat",
		Sampler:     &jaegerConf.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter:    &jaegerConf.ReporterConfig{
			LogSpans: false,
		},
	}

	jLogger := jaegerLog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegerConf.Logger(jLogger),
		jaegerConf.Metrics(jMetricsFactory),
	)

	if err != nil {
		log.Fatal("failed to initialize tracer")
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	var config config.Config

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	s := internal.NewServer()
	log.Println("waiting routes...")
	log.Fatal(http.ListenAndServe(config.Server.Port, s.Router))
}
