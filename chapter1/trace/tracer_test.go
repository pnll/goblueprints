package trace

import (
	"bytes"
	"testing"
)

// TestNew는 추적 동작을 테스트한다.
func TestNew(t *testing.T) {

	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("New에서 nil을 리턴하면 안 됩니다.")
	}
	tracer.Trace("Hello trace 패키지")
	if buf.String() != "Hello trace 패키지\n" {
		t.Errorf("Trace는 '%s'을 써서는 안 됩니다.", buf.String())
	}

}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("something")
}
