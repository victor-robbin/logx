package logx

import (
	"bytes"
	"context"
	"testing"
)

func TestFromContext(t *testing.T) {
	ctx := context.Background()
	ctx = WithRunID(ctx, "run-xyz")
	ctx = WithWorkflowID(ctx, "wf-abc")
	ctx = WithTraceID(ctx, "trace-123")

	logger := FromContext(ctx)

	if !bytes.Contains(logger.Context, []byte(`"run_id":"run-xyz"`)) {
		t.Error("missing run_id in logger context")
	}
	if !bytes.Contains(logger.Context, []byte(`"workflow_id":"wf-abc"`)) {
		t.Error("missing workflow_id in logger context")
	}
	if !bytes.Contains(logger.Context, []byte(`"trace_id":"trace-123"`)) {
		t.Error("missing trace_id in logger context")
	}
}
