package main

import (
	"context"
	"os"

	"github.com/couchbaselabs/gocbcolumnar"
	"github.com/sirupsen/logrus"
)

// #tag::loggerwrapper[]
type MyLogrusLogger struct {
	logger *logrus.Logger
}

// Log function doesn't match the gocb Log function so we need to do a bit of marshalling.
func (logger *MyLogrusLogger) Log(level cbcolumnar.LogLevel, offset int, format string, v ...interface{}) error {
	// We need to do some conversion between gocb and logrus levels as they don't match up.
	var logrusLevel logrus.Level
	switch level {
	case cbcolumnar.LogError:
		logrusLevel = logrus.ErrorLevel
	case cbcolumnar.LogWarn:
		logrusLevel = logrus.WarnLevel
	case cbcolumnar.LogInfo:
		logrusLevel = logrus.InfoLevel
	case cbcolumnar.LogDebug:
		logrusLevel = logrus.DebugLevel
	case cbcolumnar.LogTrace:
		logrusLevel = logrus.TraceLevel
	case cbcolumnar.LogSched:
		logrusLevel = logrus.TraceLevel
	case cbcolumnar.LogMaxVerbosity:
		logrusLevel = logrus.TraceLevel
	}

	// Send the data to the logrus Logf function to make sure that it gets formatted correctly.
	logger.logger.Logf(logrusLevel, format, v...)
	return nil
}

// #end::loggerwrapper[]

func main() {
	// #tag::creation[]
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)

	cbcolumnar.SetLogger(&MyLogrusLogger{
		logger: logger,
	})
	// #end::creation[]

	connStr := "couchbases://..."
	username := "..."
	password := "..."

	cluster, err := cbcolumnar.NewCluster(
		connStr,
		cbcolumnar.NewCredential(username, password),
	)
	handleErr(err)

	_, err = cluster.ExecuteQuery(context.Background(), "select 1")
	if err != nil {
		panic(err)
	}

	err = cluster.Close()
	if err != nil {
		panic(err)
	}
}
