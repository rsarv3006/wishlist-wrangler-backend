package coresharedlogger

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type LogField struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
	Fatal(msg string, fields ...LogField)
	WithContext(ctx context.Context) Logger
	WithFields(fields ...LogField) Logger
}

type LoggerFactory interface {
	CreateLogger(name string) Logger
}

type StandardLogger struct {
	logger *log.Logger
	fields []LogField
}

func NewStandardLoggerFactory() LoggerFactory {
	return &StandardLoggerFactory{}
}

type StandardLoggerFactory struct{}

func (f *StandardLoggerFactory) CreateLogger(name string) Logger {
	return &StandardLogger{
		logger: log.New(os.Stdout, fmt.Sprintf("[%s] ", name), log.LstdFlags),
	}
}

func (l *StandardLogger) log(level LogLevel, msg string, fields ...LogField) {
	allFields := append(l.fields, fields...)
	fieldStr := ""
	for _, f := range allFields {
		fieldStr += fmt.Sprintf(" %s=%v", f.Key, f.Value)
	}
	l.logger.Printf("[%s] %s%s", getLevelString(level), msg, fieldStr)
}

func (l *StandardLogger) Debug(msg string, fields ...LogField) { l.log(DebugLevel, msg, fields...) }
func (l *StandardLogger) Info(msg string, fields ...LogField)  { l.log(InfoLevel, msg, fields...) }
func (l *StandardLogger) Warn(msg string, fields ...LogField)  { l.log(WarnLevel, msg, fields...) }
func (l *StandardLogger) Error(msg string, fields ...LogField) { l.log(ErrorLevel, msg, fields...) }
func (l *StandardLogger) Fatal(msg string, fields ...LogField) {
	l.log(FatalLevel, msg, fields...)
	os.Exit(1)
}

func (l *StandardLogger) WithContext(ctx context.Context) Logger {
	return l.WithFields(LogField{Key: "trace_id", Value: ctx.Value("trace_id")})
}

func (l *StandardLogger) WithFields(fields ...LogField) Logger {
	newLogger := &StandardLogger{
		logger: l.logger,
		fields: make([]LogField, len(l.fields)+len(fields)),
	}
	copy(newLogger.fields, l.fields)
	copy(newLogger.fields[len(l.fields):], fields)
	return newLogger
}

func getLevelString(level LogLevel) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type PatternLogger struct {
	logger  *log.Logger
	fields  []LogField
	pattern string
}

type PatternLoggerFactory struct {
	pattern string
}

func NewPatternLoggerFactory(pattern string) LoggerFactory {
	return &PatternLoggerFactory{pattern: pattern}
}

func (f *PatternLoggerFactory) CreateLogger(name string) Logger {
	return &PatternLogger{
		logger:  log.New(os.Stdout, "", 0),
		pattern: f.pattern,
	}
}

func (l *PatternLogger) formatMessage(level LogLevel, msg string, fields ...LogField) string {
	allFields := append(l.fields, fields...)

	fieldMap := make(map[string]interface{})
	for _, f := range allFields {
		fieldMap[f.Key] = f.Value
	}

	traceID := "-"
	spanID := "-"
	if val, ok := fieldMap["trace_id"]; ok && val != nil {
		traceID = fmt.Sprintf("%v", val)
	}
	if val, ok := fieldMap["span_id"]; ok && val != nil {
		spanID = fmt.Sprintf("%v", val)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	levelStr := fmt.Sprintf("%-5s", getLevelString(level))

	pattern := l.pattern
	if pattern == "" {
		pattern = "%d{yyyy-MM-dd HH:mm:ss} [%X{traceId:-},%X{spanId:-}] ${LOG_LEVEL_PATTERN:-%5p} %m%n"
	}

	output := strings.ReplaceAll(pattern, "%d{yyyy-MM-dd HH:mm:ss}", timestamp)
	output = strings.ReplaceAll(output, "%X{traceId:-}", traceID)
	output = strings.ReplaceAll(output, "%X{spanId:-}", spanID)
	output = strings.ReplaceAll(output, "${LOG_LEVEL_PATTERN:-%5p}", levelStr)
	output = strings.ReplaceAll(output, "%m", msg)
	output = strings.ReplaceAll(output, "%n", "\n")

	fieldStr := ""
	for _, f := range allFields {
		if f.Key != "trace_id" && f.Key != "span_id" {
			fieldStr += fmt.Sprintf(" %s=%v", f.Key, f.Value)
		}
	}
	if fieldStr != "" {
		output = strings.TrimSuffix(output, "\n") + fieldStr + "\n"
	}

	return output
}

func (l *PatternLogger) log(level LogLevel, msg string, fields ...LogField) {
	output := l.formatMessage(level, msg, fields...)
	l.logger.Print(output)
}

func (l *PatternLogger) Debug(msg string, fields ...LogField) { l.log(DebugLevel, msg, fields...) }
func (l *PatternLogger) Info(msg string, fields ...LogField)  { l.log(InfoLevel, msg, fields...) }
func (l *PatternLogger) Warn(msg string, fields ...LogField)  { l.log(WarnLevel, msg, fields...) }
func (l *PatternLogger) Error(msg string, fields ...LogField) { l.log(ErrorLevel, msg, fields...) }
func (l *PatternLogger) Fatal(msg string, fields ...LogField) {
	l.log(FatalLevel, msg, fields...)
	os.Exit(1)
}

func (l *PatternLogger) WithContext(ctx context.Context) Logger {
	fields := []LogField{
		{Key: "trace_id", Value: ctx.Value("trace_id")},
		{Key: "span_id", Value: ctx.Value("span_id")},
	}
	return l.WithFields(fields...)
}

func (l *PatternLogger) WithFields(fields ...LogField) Logger {
	newLogger := &PatternLogger{
		logger:  l.logger,
		pattern: l.pattern,
		fields:  make([]LogField, len(l.fields)+len(fields)),
	}
	copy(newLogger.fields, l.fields)
	copy(newLogger.fields[len(l.fields):], fields)
	return newLogger
}
