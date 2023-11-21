package concourse

import (
	"github.com/sirupsen/logrus"
	"io"
)

func setupLogging(stderr io.Writer) {
	logrus.SetOutput(stderr)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}
