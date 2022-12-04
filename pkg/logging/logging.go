package logging

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

type HookWriter struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *HookWriter) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *HookWriter) Levels() []logrus.Level {
	return hook.LogLevels
}

// Получение логера
var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

// Создание логера
func Init() {
	//New - конструктор logrus, который возвращвет нам дефолтный логгер
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		//CallerPrettyfier - нам передается frame, в котором проимходит логгирование
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := f.File
			return fmt.Sprintf("%s", f.Function), fmt.Sprintf("%s : %d", filename, f.Line)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}
	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)
	//Создание Hook`a
	l.AddHook(&HookWriter{
		Writer:    []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)

}
