package logfmt

import (
	"bytes"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

func TestLogFmt(t *testing.T) {
	var (
		buf bytes.Buffer
		e   logrus.Entry
	)
	e.Buffer = &buf
	e.Time = time.Now()
	e.Message = "Hello Wolrd"
	e.Data = map[string]interface{}{
		"a": "a",
		"b": true,
		"c": 1,
	}
	DefaultFormatter.Format(&e)
	t.Log(buf.String())
}

func BenchmarkPlain(b *testing.B) {
	SortKeys = true
	var (
		f   logrus.Formatter = DefaultFormatter
		buf bytes.Buffer
		e   logrus.Entry
	)
	e.Buffer = &buf
	e.Time = time.Now()
	e.Message = "Hello Wolrd"
	for i := 0; i < b.N; i++ {
		buf.Reset()
		f.Format(&e)
	}
}

func BenchmarkLogrusText(b *testing.B) {
	var (
		f logrus.Formatter = &logrus.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02/15:04:05",
			DisableSorting:  false,
		}
		buf bytes.Buffer
		e   logrus.Entry
	)
	e.Buffer = &buf
	e.Time = time.Now()
	e.Message = "Hello Wolrd"
	for i := 0; i < b.N; i++ {
		buf.Reset()
		f.Format(&e)
	}
}

func BenchmarkLogrusJSON(b *testing.B) {
	var (
		f logrus.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02/15:04:05",
		}
		buf bytes.Buffer
		e   logrus.Entry
	)
	e.Buffer = &buf
	e.Time = time.Now()
	e.Message = "Hello Wolrd"
	for i := 0; i < b.N; i++ {
		buf.Reset()
		f.Format(&e)
	}
}
