package common

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
)

// MethonTracer methon tracer
type MethonTracer struct {
	span    opentracing.Span
	tracing opentracing.Tracer

	starting bool
}

// Begin 开始监听
func (r *MethonTracer) Begin(ctx interface{}) {

	mthonctx, ok := ctx.(context.Context)
	if !ok {
		return
	}

	parentSpan := opentracing.SpanFromContext(mthonctx)
	if parentSpan != nil {
		r.span = r.tracing.StartSpan("MethonSpan", opentracing.ChildOf(parentSpan.Context()))
	}

	r.starting = true
}

func (r *MethonTracer) SetTag(key string, val interface{}) {
	if r.span != nil {
		r.span.SetTag(key, val)
	}
}

func (r *MethonTracer) GetID() string {
	if r.span != nil {
		if sc, ok := r.span.Context().(jaeger.SpanContext); ok {
			return sc.TraceID().String()
		}
	}

	return ""
}

// End 结束监听
func (r *MethonTracer) End(ctx interface{}) {

	if !r.starting {
		return
	}

	if r.span != nil {
		r.span.Finish()
	}

}
