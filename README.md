# Introduction

This project contains *utils* for Go projects.

# Projects

## Configuration
Read configuration from file (format `key=value`).

### Functions:
```go
func LoadConfig(filename string) (map[string]string, error) {}
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

// Logger creation, from parameters
func NewLogger(output string, level LogLevel) (*Logger, error) {}
// Logger creation, from configuration file
func NewLoggerFromConfig(filename string) (*Logger, error) {}
// Don't forget to close the Logger afterward
func (logger *Logger) Close() {}

// Logging methods…
// Trace
func (logger *Logger) Trace(v ...interface{}) {}
// Trace with format
func (logger *Logger) Tracef(format string, v ...interface{}) {}
// Trace begin
func (logger *Logger) TraceBegin(f string) {}
// Trace end
func (logger *Logger) TraceEnd(f string) {}
// Debug
func (logger *Logger) Debug(v ...interface{}) {}
// Debug with format
func (logger *Logger) Debugf(format string, v ...interface{}) {}
// Info
func (logger *Logger) Info(v ...interface{}) {}
// Info with format
func (logger *Logger) Infof(format string, v ...interface{}) {}
// Warn
func (logger *Logger) Warn(v ...interface{}) {}
// Warn with format
func (logger *Logger) Warnf(format string, v ...interface{}) {}
// Error
func (logger *Logger) Error(v ...interface{}) {}
// Error with format
func (logger *Logger) Errorf(format string, v ...interface{}) {}
// Fatal
func (logger *Logger) Fatal(v ...interface{}) {}
// Fatal with format
func (logger *Logger) Fatalf(format string, v ...interface{}) {}
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

// Mailer creation, from parameters
func NewMailer(server, port, from, to string) *Mailer {}
// Mailer creation, from configuration file
func NewMailerFromConfig(filename string) (*Mailer, error) {}
// Send an email…
func (mail *Mailer) SendMail(msg string) error {}
```
