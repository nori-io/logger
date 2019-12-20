package logger_test

import (
	"testing"

	loggerNoriCommon "github.com/nori-io/nori-common/logger"
	"github.com/stretchr/testify/assert"

	"github.com/nori-io/logger"
)



func TestNew(t *testing.T) {
	// тут у тебя создается твой логгер
	a := assert.New(t)
	bufferSize := 54

	// создается твой хук
	//hook := *new(logger.Hook)
	fields := make([]loggerNoriCommon.Field, 1)
	fields = append(fields, loggerNoriCommon.Field{Key: "1", Value: "test1"})
	levelHooks := logger.LevelHooks{}
	levelHooks.Add(&TestHook{Fired:false})
	levelHooks.Fire(loggerNoriCommon.LevelInfo, fields)


	result := make([]byte, bufferSize)
	//_, err:= buf.Read(result)
	//logTest1 := logger.New()
	//logTest1.Info("test")
	result2 := make([]byte, bufferSize)
	//	_, err = buf.Read(result2)
	a.Equal(string(result), string(result2))

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

func TestHookFires(t *testing.T) {
	//hook := new(TestHook)

	/*LogAndAssertJSON(t, func(log *Logger) {
		log.Hooks.Add(hook)
		assert.Equal(t, hook.Fired, false)

		log.Print("test")
	}, func(fields Fields) {
		assert.Equal(t, hook.Fired, true)
	})*/
	/*}

	  type ModifyHook struct {
	  }*/

	/*func (hook *ModifyHook) Fire(entry *Entry) error {
	  	entry.Data["wow"] = "whale"
	  	return nil
	  }
	*/
	/*func (hook *ModifyHook) Levels() []Level {
	  	return []Level{
	  		TraceLevel,
	  		DebugLevel,
	  		InfoLevel,
	  		WarnLevel,
	  		ErrorLevel,
	  		FatalLevel,
	  		PanicLevel,
	  	}
	  }
	*/
	/*func TestHookCanModifyEntry(t *testing.T) {
	  	hook := new(ModifyHook)

	  	LogAndAssertJSON(t, func(log *Logger) {
	  		log.Hooks.Add(hook)
	  		log.WithField("wow", "elephant").Print("test")
	  	}, func(fields Fields) {
	  		assert.Equal(t, fields["wow"], "whale")
	  	})
	  }
	*/
	/*func TestCanFireMultipleHooks(t *testing.T) {
	  	hook1 := new(ModifyHook)
	  	hook2 := new(TestHook)

	  	LogAndAssertJSON(t, func(log *Logger) {
	  		log.Hooks.Add(hook1)
	  		log.Hooks.Add(hook2)

	  		log.WithField("wow", "elephant").Print("test")
	  	}, func(fields Fields) {
	  		assert.Equal(t, fields["wow"], "whale")
	  		assert.Equal(t, hook2.Fired, true)
	  	})
	  }
	*/

	/*type SingleLevelModifyHook struct {
		ModifyHook
	}*/

	/*func (h *SingleLevelModifyHook) Levels() []Level {
	  	return []Level{InfoLevel}
	  }
	*/
	/*func TestHookEntryIsPristine(t *testing.T) {
	  	l := New()
	  	b := &bytes.Buffer{}
	  	l.Formatter = &JSONFormatter{}
	  	l.Out = b
	  	l.AddHook(&SingleLevelModifyHook{})

	  	l.Error("error message")
	  	data := map[string]string{}
	  	err := json.Unmarshal(b.Bytes(), &data)
	  	require.NoError(t, err)
	  	_, ok := data["wow"]
	  	require.False(t, ok)
	  	b.Reset()

	  	l.Info("error message")
	  	data = map[string]string{}
	  	err = json.Unmarshal(b.Bytes(), &data)
	  	require.NoError(t, err)
	  	_, ok = data["wow"]
	  	require.True(t, ok)
	  	b.Reset()

	  	l.Error("error message")
	  	data = map[string]string{}
	  	err = json.Unmarshal(b.Bytes(), &data)
	  	require.NoError(t, err)
	  	_, ok = data["wow"]
	  	require.False(t, ok)
	  	b.Reset()
	  }
	*/
	/*type ErrorHook struct {
	  	Fired bool
	  }
	*/
	/*func (hook *ErrorHook) Fire(entry *Entry) error {
	  	hook.Fired = true
	  	return nil
	  }
	*/
	/*func (hook *ErrorHook) Levels() []Level {
	  	return []Level{
	  		ErrorLevel,
	  	}
	  }
	*/
	/*func TestErrorHookShouldntFireOnInfo(t *testing.T) {
	  	hook := new(ErrorHook)

	  	LogAndAssertJSON(t, func(log *Logger) {
	  		log.Hooks.Add(hook)
	  		log.Info("test")
	  	}, func(fields Fields) {
	  		assert.Equal(t, hook.Fired, false)
	  	})
	  }
	*/
	/*func TestErrorHookShouldFireOnError(t *testing.T) {
	  	hook := new(ErrorHook)

	  	LogAndAssertJSON(t, func(log *Logger) {
	  		log.Hooks.Add(hook)
	  		log.Error("test")
	  	}, func(fields Fields) {
	  		assert.Equal(t, hook.Fired, true)
	  	})
	  }
	*/
	/*func TestAddHookRace(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(2)
		hook := new(ErrorHook)
		LogAndAssertJSON(t, func(log *Logger) {
			go func() {
				defer wg.Done()
				log.AddHook(hook)
			}()
			go func() {
				defer wg.Done()
				log.Error("test")
			}()
			wg.Wait()
		}, func(fields Fields) {
			// the line may have been logged
			// before the hook was added, so we can't
			// actually assert on the hook
		})
	}*/
}
