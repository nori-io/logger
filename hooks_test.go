package logger_test

import (
	"bytes"
	"log/syslog"
	"testing"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"

	"github.com/nori-io/logger"
)

func TestLocalhostAddAndPrint(t *testing.T) {

	buf := bytes.Buffer{}

	hook, err := logger.NewFileHook("test_file")

	log := logger.New(logger.SetJsonFormatter(), logger.SetOutWriter(&buf), logger.SetFileHook(*hook))

	if err != nil {
		t.Errorf("Can't create hook")
	}

	log.Info("done")

}

type TestHook struct {
	Fired bool
}

func (hook *TestHook) Fire(fields []loggerNoriCommon.Field) error {
	hook.Fired = true
	return nil
}

func (hook *TestHook) Levels() []loggerNoriCommon.Level {
	return []loggerNoriCommon.Level{
		loggerNoriCommon.LevelFatal,
		loggerNoriCommon.LevelPanic,
		loggerNoriCommon.LevelNotice,
		loggerNoriCommon.LevelCritical,
		loggerNoriCommon.LevelError,
		loggerNoriCommon.LevelWarning,
		loggerNoriCommon.LevelInfo,
		loggerNoriCommon.LevelDebug,
	}
}

type SyslogHook struct {
	Writer        *syslog.Writer
	SyslogNetwork string
	SyslogRaddr   string
}

/*func (hook *FileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		return hook.Writer.Crit(line)
	case logrus.FatalLevel:
		return hook.Writer.Crit(line)
	case logrus.ErrorLevel:
		return hook.Writer.Err(line)
	case logrus.WarnLevel:
		return hook.Writer.Warning(line)
	case logrus.InfoLevel:
		return hook.Writer.Info(line)
	case logrus.DebugLevel, logrus.TraceLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}*/
