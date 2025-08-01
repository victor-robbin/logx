package main

import (
	"context"
	"time"

	"github.com/victor-robbin/logx"
)

func main() {
	logx.Init(logx.InitConfig{
		Service:     "demo-cli",
		Environment: "local",
		Version:     "v0.1.0",
		Level:       "debug",
		LogToFile:   true,
		LogPath:     "./logs/demo-cli.log",
		MaxSizeMB:   5,
		MaxBackups:  2,
	})

	logx.Info("cli started", map[string]interface{}{
		"ts": time.Now().Format(time.RFC3339),
	})

	// Simulate contextual usage
	ctx := context.Background()
	ctx = logx.WithRunID(ctx, "run-xyz")
	ctx = logx.WithWorkflowID(ctx, "wf-abc")
	ctx = logx.WithTraceID(ctx, "trace-123")

	logx.InfoCtx(ctx, "processing workflow", map[string]interface{}{
		"step": "initial",
	})

	logx.ErrorCtx(ctx, "failed to connect", map[string]interface{}{
		"retries": 3,
	})
}
