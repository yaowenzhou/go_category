package configs

import (
	"fmt"
	"go_category/consts"
	"os"

	"github.com/opentracing/opentracing-go"
)

func init() {
	var err error
	LogFile, err = os.OpenFile(consts.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		panic(fmt.Errorf("open config file: %s", err.Error()))
	}
	err = initTracer()
	if err != nil {
		panic(fmt.Errorf("init tracer fail: %s", err.Error()))
	}
	opentracing.SetGlobalTracer(Tracer)
}
