package configs

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

var Tracer opentracing.Tracer
var Closer io.Closer

// 创建jaeger链路追踪实例
func initTracer() (err error) {
	cfg := &jaegercfg.Configuration{
		ServiceName: "category",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "192.168.1.99:6831",
		},
	}
	Tracer, Closer, err = cfg.NewTracer()
	return err
}

// CloseIOCloser 退出时关闭 Closer
func CloseIOCloser() {
	Closer.Close()
}
