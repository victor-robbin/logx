# logx

**Reusable structured logger for Go microservices**, built on top of [`phuslu/log`](https://github.com/phuslu/log).

## ✨ Features

- ✅ Minimal setup
- 📁 File + console output support
- 📦 Injects service metadata (`service`, `env`, `version`)
- 🧠 Context-aware fields (`trace_id`, `workflow_id`, `run_id`)
- 🧪 Simple, testable, and dependency-light

---

## 🚀 Quick Start

```go
import (
  "context"
  "github.com/victor-robbin/logx"
)

func main() {
  logx.Init(logx.InitConfig{
    Service:     "worker",
    Environment: "prod",
    Version:     "v1.0.0",
    Level:       "info",
    LogToFile:   true,
    LogPath:     "/var/log/worker.log",
    MaxSizeMB:   10,
    MaxBackups:  5,
  })

  ctx := logx.WithTraceID(context.Background(), "trace-123")
  logx.InfoCtx(ctx, "task started", map[string]interface{}{"task": "email"})
}
```

---

## 🧩 Context Fields

Supports contextual injection of tracing fields:

- `trace_id`
- `workflow_id`
- `run_id`

Example:

```go
ctx := context.Background()
ctx = logx.WithRunID(ctx, "run-xyz")
ctx = logx.WithWorkflowID(ctx, "wf-abc")
ctx = logx.WithTraceID(ctx, "trace-123")

logx.ErrorCtx(ctx, "workflow failed", map[string]interface{}{"retries": 3})
```

---

## 🧪 Testing Helpers

Use `TestWriter` to capture and assert log output in tests:

```go
buf := &bytes.Buffer{}
testWriter := &logx.TestWriter{Buf: buf}

logger := logx.FromContext(ctx)
l := *logger
l.Writer = testWriter

l.Info().Str("step", "test").Msg("test case")

// assert buf.String() contains expected fields
```

---

## 📂 Example CLI App

See [examples/cli/main.go](examples/cli/main.go) for CLI demo.

---

## 🔒 License

MIT
