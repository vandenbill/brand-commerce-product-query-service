package util

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"os"
)

func ConfigureJaeger() io.Closer {
	jaegerHOSTPORT := os.Getenv("JAEGER_HOST_PORT")
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 10,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerHOSTPORT,
		},
	}
	closer, err := cfg.InitGlobalTracer(
		"product-query-service",
	)
	if err != nil {
		log.Panicf("Could not initialize jaeger tracer: %s", err.Error())
		return nil
	}

	return closer
}
