package logger

import (
  "github.com/sirupsen/logrus"
  "os"
)

func Logger() *logrus.Logger {
  var log = logrus.New()
  log.Out = os.Stdout
  log.Formatter = &logrus.JSONFormatter{}
  log.Level = logrus.DebugLevel
  return log
}
