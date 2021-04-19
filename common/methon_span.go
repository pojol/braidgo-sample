package common

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/pojol/braid-go/module/tracer"
)

// MethonTracer methon tracer
type MethonTracer struct {
	span    opentracing.Span
	tracing opentracing.Tracer

	starting bool
}

func CreateMethonSpanFactory() tracer.SpanFactory {
	return func(tracing interface{}) (tracer.ISpan, error) {

		t, ok := tracing.(opentracing.Tracer)
		if !ok {
			return nil, errors.New("")
		}

		rt := &MethonTracer{
			tracing: t,
		}

		return rt, nil
	}
}

// Begin 开始监听
func (r *MethonTracer) Begin(ctx interface{}, tags ...tracer.SpanTag) {

	mthonctx, ok := ctx.(context.Context)
	if !ok {
		return
	}

	parentSpan := opentracing.SpanFromContext(mthonctx)
	if parentSpan != nil {
		r.span = r.tracing.StartSpan("MethonSpan", opentracing.ChildOf(parentSpan.Context()))
		for _, v := range tags {
			r.span.SetTag(v.Key, v.Val)
		}
	}

	r.starting = true
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
