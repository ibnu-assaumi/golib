package golib

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelString(t *testing.T) {

	t.Parallel()

	var l Level

	t.Run("TraceLevel String", func(t *testing.T) {
		l = TraceLevel
		assert.Equal(t, "trace", l.String())
	})

	t.Run("DebugLevel String", func(t *testing.T) {
		l = DebugLevel
		assert.Equal(t, "debug", l.String())
	})

	t.Run("InfoLevel String", func(t *testing.T) {
		l = InfoLevel
		assert.Equal(t, "info", l.String())
	})

	t.Run("WarnLevel String", func(t *testing.T) {
		l = WarnLevel
		assert.Equal(t, "warning", l.String())
	})

	t.Run("ErrorLevel String", func(t *testing.T) {
		l = ErrorLevel
		assert.Equal(t, "error", l.String())
	})

	t.Run("FatalLevel String", func(t *testing.T) {
		l = FatalLevel
		assert.Equal(t, "fatal", l.String())
	})

	t.Run("PanicLevel String", func(t *testing.T) {
		l = PanicLevel
		assert.Equal(t, "panic", l.String())
	})

	t.Run("Unknown String", func(t *testing.T) {
		l = 9999
		assert.Equal(t, "unknown", l.String())
	})
}

func TestInitLogger(t *testing.T) {
	t.Run("InitLogger", func(t *testing.T) {
		InitLogger("test", "test", "test")
		assert.Equal(t, "test", TOPIC)
		assert.Equal(t, "test", LogTag)
		assert.Equal(t, "test", Env)
	})
}

func TestLogContext(t *testing.T) {
	c := "test"
	s := "test"
	customTags := make([]map[string]interface{}, 1)
	customTag := make(map[string]interface{})

	t.Run("SUCCESS LOGCONTEXT", func(t *testing.T) {
		customTag["test"] = "test"
		customTags = append(customTags, customTag)
		assert.NotNil(t, LogContext(c, s, customTags))
	})
}

var (
	testConst = "any message"
	testTags  = map[string]interface{}{"test": testConst}
)
var testCasesLog = []struct {
	name    string
	level   Level
	message string
	context string
	scope   string
	tags    map[string]interface{}
}{
	{
		name:    "#1 Debug",
		level:   DebugLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#2 Info",
		level:   InfoLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#3 Warn",
		level:   WarnLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#4 Error",
		level:   ErrorLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
	{
		name:    "#5 Panic",
		level:   PanicLevel,
		message: testConst,
		context: testConst,
		scope:   testConst,
		tags:    testTags,
	},
}

func TestLog(t *testing.T) {

	t.Parallel()

	for _, tc := range testCasesLog {
		t.Run(tc.name, func(*testing.T) {
			Log(tc.level, tc.message, tc.context, tc.scope, tc.tags)
		})
	}
}

func TestLogError(t *testing.T) {
	err := errors.New("test")
	ctx := "test"

	t.Parallel()

	t.Run("LOG ERROR", func(*testing.T) {
		msg := "test"
		LogError(err, ctx, msg)
	})
}

func TestNewFileResultLogger(t *testing.T) {

	t.Run("SUCCESS newFileResultLogger", func(t *testing.T) {
		re := regexp.MustCompile(`^(.*golib)`)
		cwd, _ := os.Getwd()
		rootPath := string(re.Find([]byte(cwd)))
		res := newFileResultLogger(rootPath)
		assert.Equal(t, rootPath, res.baseDir)
	})

	t.Run("ERROR newFileResultLogger", func(t *testing.T) {
		re := regexp.MustCompile(`^(.*gox)`)
		cwd, _ := os.Getwd()
		rootPath := string(re.Find([]byte(cwd)))
		res := newFileResultLogger(rootPath)
		assert.Equal(t, rootPath, res.baseDir)
	})
}

func TestFileResultLoggerLastError(t *testing.T) {
	t.Run("NIL LastError", func(t *testing.T) {
		f := &FileResultLogger{}
		assert.Nil(t, f.LastError())
	})

	t.Run("NIL LastError", func(t *testing.T) {
		f := &FileResultLogger{
			lastError: errors.New("error"),
		}
		assert.Error(t, f.LastError())
	})
}

func TestFileResultLoggerGetFileName(t *testing.T) {
	t.Run("SUCCESS GetFileName", func(t *testing.T) {
		f := &FileResultLogger{}
		assert.NotEqual(t, "", f.GetFileName("test"))
	})
}

func TestFileResultLoggerGet(t *testing.T) {
	t.Run("ERROR Get", func(t *testing.T) {
		f := &FileResultLogger{}
		assert.Equal(t, "", f.Get(""))
	})
}

func TestFileResultLoggerStore(t *testing.T) {
	t.Run("SUCCESS Store", func(t *testing.T) {
		re := regexp.MustCompile(`^(.*golib)`)
		cwd, _ := os.Getwd()
		rootPath := string(re.Find([]byte(cwd)))
		f := &FileResultLogger{}
		f.baseDir = rootPath
		s := f.Store("go.sum", []byte("test"))
		assert.NotEqual(t, "", s)
	})

	t.Run("ERROR Store", func(t *testing.T) {
		f := &FileResultLogger{}
		s := f.Store("test", []byte("test"))
		assert.NotEqual(t, "", s)
	})
}

func TestFileResultLoggerRequestResponse(t *testing.T) {
	t.Run("ERROR RequestResponse", func(t *testing.T) {
		f := &FileResultLogger{}
		s := f.RequestResponse("test", "test")
		assert.Equal(t, "", s)
	})
}

func TestGetResultLogger(t *testing.T) {
	t.Run("NULL BASE GetResultLogger", func(t *testing.T) {
		assert.NotNil(t, GetResultLogger())
	})
}

func TestStoreRequestResponse(t *testing.T) {
	t.Run("", func(*testing.T) {
		s := StoreRequestResponse("200", []byte("test"), []byte("test"))
		fmt.Println(s)
	})
}
