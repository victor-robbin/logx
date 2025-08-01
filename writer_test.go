package logx

import (
	"bytes"

	"github.com/phuslu/log"
)

type TestWriter struct {
	Buf *bytes.Buffer
}

func (tw *TestWriter) WriteEntry(e *log.Entry) (int, error) {
	// Use built-in ConsoleWriter to format Entry into our buffer
	cw := &log.ConsoleWriter{
		Writer:         tw.Buf,
		ColorOutput:    false,
		QuoteString:    false,
		EndWithMessage: true,
	}
	return cw.WriteEntry(e)
}
