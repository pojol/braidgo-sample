package middleware

import (
	"github.com/labstack/echo/v4"
	emid "github.com/labstack/echo/v4/middleware"
	"github.com/pojol/braid"
	"github.com/pojol/braid/module/tracer"
	"github.com/pojol/braid/modules/jaegertracing"
)

type (
	// ReqTraceConfig 请求追踪中间件配置项
	ReqTraceConfig struct {
		Skipper emid.Skipper
	}
)

var (
	// DefaultReqTraceConfig 默认请求限制器配置
	defaultReqTraceConfig = ReqTraceConfig{
		Skipper: emid.DefaultSkipper,
	}

	qps uint64 = 0
)

// ReqTrace 获取基于默认配置的trace
func ReqTrace() echo.MiddlewareFunc {
	/*
		go func() {
			for {
				fmt.Println("qps", qps)
				atomic.SwapUint64(&qps, 0)
				time.Sleep(time.Second)
			}
		}()
	*/
	return ReqTraceWithConfig(defaultReqTraceConfig)
}

// ReqTraceWithConfig 基于配置生成trace
func ReqTraceWithConfig(cfg ReqTraceConfig) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if cfg.Skipper(c) {
				return next(c)
			}

			tracing := braid.Tracer()
			var span tracer.ISpan

			if tracing != nil {
				span, err = tracing.GetSpan(jaegertracing.EchoSpan)
				if err != nil {
					return next(c)
				}
			}
			span.Begin(c)

			defer func() {
				span.End(c)
				//atomic.AddUint64(&qps, 1)
			}()

			return next(c)
		}
	}
}
