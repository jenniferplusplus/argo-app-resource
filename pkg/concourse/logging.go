package concourse

import (
	"github.com/sirupsen/logrus"
	"io"
)

func setupLogging(stderr io.Writer, debug bool) {
	logrus.SetOutput(stderr)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
