# Introduction

This project contains *utils* for Go projects.

# Projects

## Configuration
Read configuration from file (format `key=value`).

### Functions:
```go
func LoadConfig(filename string) (map[string]string, error)
```

## Logging
Ease logging usage.

### Functions:
```go
// Logging level type
type LogLevel int
const (
    LEVEL_TRACE LogLevel
    LEVEL_DEBUG
    LEVEL_INFO
    LEVEL_WARN
    LEVEL_ERROR
    LEVEL_OFF
)

// Logger type, is a log.Logger
type Logger struct {
    *log.Logger

    Level  LogLevel
    Output *os.File
}

// Logger creation, from parameters or configuration file
func NewLogger(output string, level LogLevel) (*Logger, error)
func NewLoggerFromConfig(filename string) (*Logger, error)
// Don't forget to close the Logger afterward
func (logger *Logger) Close()

// Logging methods…
func (logger *Logger) Trace(v ...interface{})
func (logger *Logger) Tracef(format string, v ...interface{})
func (logger *Logger) TraceBegin(f string)
func (logger *Logger) TraceEnd(f string)
func (logger *Logger) Debug(v ...interface{})
func (logger *Logger) Debugf(format string, v ...interface{})
func (logger *Logger) Info(v ...interface{})
func (logger *Logger) Infof(format string, v ...interface{})
func (logger *Logger) Warn(v ...interface{})
func (logger *Logger) Warnf(format string, v ...interface{})
func (logger *Logger) Error(v ...interface{})
func (logger *Logger) Errorf(format string, v ...interface{})
func (logger *Logger) Fatal(v ...interface{})
func (logger *Logger) Fatalf(format string, v ...interface{})
```

## Mailer
Easily send emails.

### Functions:
```go
// Mailer type
type Mailer struct {
    Server string
    Port   string
    From   string
    To     string
}

// Mailer creation, from parameters or configuration file
func NewMailer(server, port, from, to string) *Mailer
func NewMailerFromConfig(filename string) (*Mailer, error)
// Send an email…
func (mail *Mailer) SendMail(msg string) error
```
