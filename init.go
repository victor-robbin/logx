package logx

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/phuslu/log"
)

var (
	logger   log.Logger
	initOnce sync.Once
)

type InitConfig struct {
	Service     string
	Environment string
	Version     string
	Level       string

	LogToFile  bool
	LogPath    string
	MaxSizeMB  int64
	MaxBackups int
	MaxAgeDays int // for converting to time.Duration
	Compress   bool
}

func ResetLoggerForTest() {
	logger = log.Logger{}
	initOnce = sync.Once{}
}

// Init sets up the global logger
func Init(cfg InitConfig) {
	initOnce.Do(func() {
		level := log.ParseLevel(cfg.Level)

		// Create file writer if enabled
		var fileWriter log.Writer
		if cfg.LogToFile && cfg.LogPath != "" {
			if err := os.MkdirAll(filepath.Dir(cfg.LogPath), 0755); err != nil {
				panic("failed to create log directory: " + err.Error())
			}

			fileWriter = &log.FileWriter{
				Filename:     cfg.LogPath,
				FileMode:     0644,
				MaxSize:      cfg.MaxSizeMB * 1024 * 1024,
				MaxBackups:   cfg.MaxBackups,
				LocalTime:    true,
				EnsureFolder: true,
			}
		}

		consoleWriter := &log.ConsoleWriter{
			ColorOutput:    true,
			QuoteString:    true,
			EndWithMessage: true,
		}

		// Use MultiLevelWriter to route based on level
		writer := &log.MultiLevelWriter{
			InfoWriter:    fileWriter,
			WarnWriter:    fileWriter,
			ErrorWriter:   fileWriter,
			ConsoleWriter: consoleWriter,
			ConsoleLevel:  level, // mirror console from this level up
		}

		logger = log.Logger{
			Level:      level,
			Caller:     1,
			TimeField:  "ts",
			TimeFormat: time.RFC3339Nano,
			Writer:     writer,
		}
		logger.Context = log.NewContext(nil).
			Str("service", cfg.Service).
			Str("env", cfg.Environment).
			Str("version", cfg.Version).
			Value()
	})
}

// FromContext returns a contextual logger with run_id, workflow_id, etc.
func FromContext(ctx context.Context) *log.Logger {
	sub := logger
	fields := log.NewContext(nil)

	if v := ctx.Value(ctxKeyRunID{}); v != nil {
		fields = fields.Str("run_id", v.(string))
	}
	if v := ctx.Value(ctxKeyWorkflowID{}); v != nil {
		fields = fields.Str("workflow_id", v.(string))
	}
	if v := ctx.Value(ctxKeyTraceID{}); v != nil {
		fields = fields.Str("trace_id", v.(string))
	}

	sub.Context = fields.Value()
	return &sub
}

func Logger() *log.Logger {
	return &logger
}

// Context key types
type ctxKeyRunID struct{}
type ctxKeyWorkflowID struct{}
type ctxKeyTraceID struct{}

func WithRunID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKeyRunID{}, id)
}

func WithWorkflowID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKeyWorkflowID{}, id)
}

func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxKeyTraceID{}, id)
}
