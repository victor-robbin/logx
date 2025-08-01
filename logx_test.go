package logx_test

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/victor-robbin/logx"
)

func waitForFile(path string, retries int, delay time.Duration) bool {
	for i := 0; i < retries; i++ {
		if _, err := os.Stat(path); err == nil {
			return true
		}
		time.Sleep(delay)
	}
	return false
}
func TestInitLogger(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")

	logx.Init(logx.InitConfig{
		Service:     "test-service",
		Environment: "test",
		Version:     "v0.0.1",
		Level:       "debug",
		LogToFile:   true,
		LogPath:     logFile,
		MaxSizeMB:   1,
		MaxBackups:  1,
	})

	logger := logx.Logger()
	logger.Debug().Str("unit", "test").Msg("debug test")
	logger.Info().Msg("info test")

	if !waitForFile(logFile, 10, 20*time.Millisecond) {
		t.Fatalf("Expected log file %s to exist", logFile)
	}
}

func TestContextInjection(t *testing.T) {
	ctx := context.Background()
	ctx = logx.WithRunID(ctx, "run-1")
	ctx = logx.WithWorkflowID(ctx, "wf-2")
	ctx = logx.WithTraceID(ctx, "trace-3")

	logger := logx.FromContext(ctx)

	buf := &bytes.Buffer{}
	testWriter := &logx.TestWriter{Buf: buf}

	// Copy logger and override Writer
	l := *logger
	l.Writer = testWriter

	l.Info().Str("step", "test").Msg("context test")

	output := buf.String()
	if !strings.Contains(output, `run_id=run-1`) {
		t.Error("missing run_id")
	}
	if !strings.Contains(output, `workflow_id=wf-2`) {
		t.Error("missing workflow_id")
	}
	if !strings.Contains(output, `trace_id=trace-3`) {
		t.Error("missing trace_id")
	}
}

func TestLoggerHelpers(t *testing.T) {
	logx.ResetLoggerForTest() // <- IMPORTANT!

	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "helper.log")

	logx.Init(logx.InitConfig{
		Service:     "helper-test",
		Environment: "test",
		Version:     "v0.0.2",
		Level:       "debug",
		LogToFile:   true,
		LogPath:     logFile,
		MaxSizeMB:   1,
		MaxBackups:  1,
	})

	ctx := logx.WithTraceID(context.Background(), "trace-xyz")

	logx.Info("test info", map[string]interface{}{"x": 1})
	logx.WarnCtx(ctx, "test warn ctx", map[string]interface{}{"warn": true})
	logx.Error("test error", map[string]interface{}{"err": "value"})

	if !waitForFile(logFile, 10, 20*time.Millisecond) {
		t.Fatalf("Expected helper log file %s to exist", logFile)
	}
}
