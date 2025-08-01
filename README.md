# logx

Reusable structured logger for Go microservices.

* ✅ Powered by [phuslu/log](https://github.com/phuslu/log)
* ✅ Console + file logging
* ✅ Injects `service`, `env`, `version`, and contextual fields (e.g. `trace_id`)
* ✅ Context-aware logging for workflows and background jobs

---

## 📦 Installation

```bash
go get github.com/victor-robbin/logx@v0.1.0
# or always latest
go get github.com/victor-robbin/logx@latest
```

---

## 🚀 Quick Start

### Console-only (no file logging)

```go
logx.Init(logx.InitConfig{
  Service:     "worker",
  Environment: "dev",
  Version:     "v1.0.0",
  Level:       "debug",
  LogToFile:   false,
})
```

### With file logging

```go
logx.Init(logx.InitConfig{
  Service:     "worker",
  Environment: "prod",
  Version:     "v1.0.0",
  Level:       "info",
  LogToFile:   true,
  LogPath:     "/var/log/worker.log",
  MaxSizeMB:   5,
  MaxBackups:  3,
})
```

---

## 🧠 Contextual Logging (Trace ID, Workflow ID)

```go
ctx := context.Background()
ctx = logx.WithTraceID(ctx, "trace-abc")
ctx = logx.WithRunID(ctx, "run-123")
ctx = logx.WithWorkflowID(ctx, "wf-xyz")

logx.InfoCtx(ctx, "workflow started", map[string]any{
  "step": "init",
})
```

---

## 🧪 Sample Output

```txt
2025-08-01T12:00:00Z INF main.go:12 > service="worker" env="prod" version="v1.0.0" run_id="run-123" workflow_id="wf-xyz" trace_id="trace-abc" step="init" workflow started
```

---

## 🔫 Testing

```bash
go test ./...
```

---

## 📂 File Rotation (if enabled)

* `MaxSizeMB`: file will rotate when it exceeds N megabytes
* `MaxBackups`: number of rotated files to keep
* Logs are always written in local time if file logging is used

---

## 📌 License

MIT
