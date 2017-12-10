package anaconda

import (
	"log"
	"os"
)

// The Logger interface provides optional logging ability for the streaming API.
// It can also be used to log the rate limiting headers if desired.
type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	// Log functions
	Critical(args ...interface{})
	Criticalf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args ...interface{})

	Notice(args ...interface{})
	Noticef(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

// SetLogger sets the Logger used by the API client.
// The default logger is silent. BasicLogger will log to STDERR
// using the log package from the standard library.
func (c *TwitterApi) SetLogger(l Logger) {
	c.Log = l
}

type silentLogger struct {
}

func (_ silentLogger) Fatal(_ ...interface{})                 {}
func (_ silentLogger) Fatalf(_ string, _ ...interface{})      {}
func (_ silentLogger) Panic(_ ...interface{})                 {}
func (_ silentLogger) Panicf(_ string, _ ...interface{})      {}
func (_ silentLogger) Critical(_ ...interface{})              {}
func (_ silentLogger) Criticalf(_ string, _ ...interface{})   {}
func (_ silentLogger) Error(_ ...interface{})                 {}
func (_ silentLogger) Errorf(_ string, _ ...interface{})      {}
func (_ silentLogger) Warning(_ ...interface{})               {}
func (_ silentLogger) Warningf(_ string, _ ...interface{})    {}
func (_ silentLogger) Notice(_ ...interface{})                {}
func (_ silentLogger) Noticef(_ string, _ ...interface{})     {}
func (_ silentLogger) Info(_ ...interface{})                  {}
func (_ silentLogger) Infof(_ string, _ ...interface{})       {}
func (_ silentLogger) Debug(_ ...interface{})                 {}
func (_ silentLogger) Debugf(format string, _ ...interface{}) {}

// BasicLogger is the equivalent of using log from the standard
// library to print to STDERR.
var BasicLogger Logger

type basicLogger struct {
	log *log.Logger //func New(out io.Writer, prefix string, flag int) *Logger
}

func init() {
	BasicLogger = &basicLogger{log: log.New(os.Stderr, log.Prefix(), log.LstdFlags)}
}

func (l basicLogger) Fatal(items ...interface{})               { l.log.Fatal(items) }
func (l basicLogger) Fatalf(s string, items ...interface{})    { l.log.Fatalf(s, items) }
func (l basicLogger) Panic(items ...interface{})               { l.log.Panic(items) }
func (l basicLogger) Panicf(s string, items ...interface{})    { l.log.Panicf(s, items) }
func (l basicLogger) Critical(items ...interface{})            { l.log.Print(items) }
func (l basicLogger) Criticalf(s string, items ...interface{}) { l.log.Printf(s, items) }
func (l basicLogger) Error(items ...interface{})               { l.log.Print(items) }
func (l basicLogger) Errorf(s string, items ...interface{})    { l.log.Printf(s, items) }
func (l basicLogger) Warning(items ...interface{})             { l.log.Print(items) }
func (l basicLogger) Warningf(s string, items ...interface{})  { l.log.Printf(s, items) }
func (l basicLogger) Notice(items ...interface{})              { l.log.Print(items) }
func (l basicLogger) Noticef(s string, items ...interface{})   { l.log.Printf(s, items) }
func (l basicLogger) Info(items ...interface{})                { l.log.Print(items) }
func (l basicLogger) Infof(s string, items ...interface{})     { l.log.Printf(s, items) }
func (l basicLogger) Debug(items ...interface{})               { l.log.Print(items) }
func (l basicLogger) Debugf(s string, items ...interface{})    { l.log.Printf(s, items) }
