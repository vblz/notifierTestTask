package notifier

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	buff := bytes.NewBufferString("")
	l := funcLogger(func(format string, args ...interface{}) {
		fmt.Fprintf(buff, format, args...)
	})

	l.Logf("format %d string %s", 1, "test")
	assert.Equal(t, "format 1 string test", buff.String())
}

func TestNoOpLogger(t *testing.T) {
	buff := bytes.NewBufferString("")
	log.SetOutput(buff)
	defer log.SetOutput(os.Stdout)

	noOpLogger.Logf("test output")
	assert.Equal(t, "", buff.String())
}

func TestLeveledLogger(t *testing.T) {
	buff := bytes.NewBufferString("")
	l := funcLogger(func(format string, args ...interface{}) {
		fmt.Fprintf(buff, format, args...)
	})

	ll := toLeveledLogger(l)

	ll.Debug("test %d", 1)
	assert.Equal(t, "DEBUG test 1", buff.String())

	buff.Reset()
	ll.Info("test %d", 2)
	assert.Equal(t, "INFO test 2", buff.String())

	buff.Reset()
	ll.Warn("test %d", 3)
	assert.Equal(t, "WARN test 3", buff.String())

	buff.Reset()
	ll.Error("test %d", 4)
	assert.Equal(t, "ERROR test 4", buff.String())
}
