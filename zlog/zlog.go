package zlog

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	// Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
	Logger = &MyLogger{Logger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})}

	fw *fileWriter
)

type MyLogger struct {
	Logger zerolog.Logger
}

type LogEvent struct {
	*zerolog.Event
}

func newConsoleWriter() *zerolog.ConsoleWriter {
	cw := &zerolog.ConsoleWriter{Out: os.Stdout}
	cw.TimeFormat = time.DateTime + ".000"
	return cw
}

func Flush() {
	flushAllLog()
}

// InitLog 初始化日志
func InitLog(root, logLevel string) {
	if fw != nil {
		return
	}
	if len(root) > 0 {
		os.MkdirAll(root, 0755|0750)
	}
	fw = &fileWriter{logRoot: root}
	//  0 debug  1 info
	zerolog.TimeFieldFormat = time.DateTime + ".000"
	if root == "" {
		Logger = &MyLogger{Logger: zerolog.New(newConsoleWriter()).With().Timestamp().Caller().Logger()}
	} else {
		w := zerolog.MultiLevelWriter(fw, newConsoleWriter())
		Logger = &MyLogger{Logger: zerolog.New(w).With().Timestamp().Caller().Logger()}
	}
	l, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		l = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(l)
	addFlusher(fw)

	startFlusher()
}

func Trace() *LogEvent {
	return Logger.Trace()
}

func Debug() *LogEvent {
	return Logger.Debug()
}

func Info() *LogEvent {
	return Logger.Info()
}

func Warn() *LogEvent {
	return Logger.Warn()
}

func Error(err error) *LogEvent {
	return Logger.Error(err)
}

func Fatal() *LogEvent {
	return Logger.Fatal()
}

func Log() *LogEvent {
	return Logger.Log()
}

func (ld *LogEvent) PID(pid int64) *LogEvent {
	ld.Int64("player_id", pid)
	return ld
}

// server_id(表示本服务的id)已经被占用了
func (ld *LogEvent) SID(sid int64) *LogEvent {
	ld.Int64("sid", sid)
	return ld
}

func (ld *LogEvent) GameID(id int64) *LogEvent {
	ld.Int64("game_id", id)
	return ld
}

func (ld *LogEvent) PName(name string) *LogEvent {
	ld.Str("player_name", name)
	return ld
}

func (ld *LogEvent) AID(aid int64) *LogEvent {
	ld.Int64("account_id", aid)
	return ld
}

func (ml *MyLogger) Trace() *LogEvent {
	return &LogEvent{Event: ml.Logger.Trace()}
}

func (ml *MyLogger) Debug() *LogEvent {
	return &LogEvent{Event: ml.Logger.Debug()}
}

func (ml *MyLogger) Info() *LogEvent {
	return &LogEvent{Event: ml.Logger.Info()}
}

func (ml *MyLogger) Warn() *LogEvent {
	return &LogEvent{Event: ml.Logger.Warn()}
}

func (ml *MyLogger) Error(err error) *LogEvent {
	return &LogEvent{Event: ml.Logger.Error().Err(err)}
}

func (ml *MyLogger) Fatal() *LogEvent {
	return &LogEvent{Event: ml.Logger.Panic()}
}

func (ml *MyLogger) Log() *LogEvent {
	return &LogEvent{Event: ml.Logger.Log()}
}

func (ml *MyLogger) WithLevel(lv zerolog.Level) *LogEvent {
	return &LogEvent{Event: ml.Logger.WithLevel(lv)}
}
