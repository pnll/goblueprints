package trace

import (
	"fmt"
	"io"
)

// Tracer는 코드 전체에서 이벤트를 추적할 수 있는
// 객체를 설명하는(기술하는) 인터페이스이다.
type Tracer interface {
	Trace(...interface{})
}

// New는 지정된 io.Writer에 출력을 작성하는 새로운 Tracer를 생성한다.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

// tracer는 io.Writer에 쓰는 Tracer이다.
type tracer struct {
	out io.Writer
}

// Trace는 이 Tracers io.Writer에 인수를 쓴다.
func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// nilTracer
type nilTracer struct{}

// nil tracer에 대한 추적은 아무것도 하지 않는다.
func (t *nilTracer) Trace(a ...interface{}) {}

// Off는 Trace에 대한 호출을 무시할 Tracer를 생성한다.
func Off() Tracer {
	return &nilTracer{}
}
