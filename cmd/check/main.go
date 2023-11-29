package main

import (
	"argo-app-resource/pkg/concourse"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	command := concourse.NewTask(
		os.Stdin,
		os.Stderr,
		os.Stdout,
		os.Args,
	)

	err := command.Check()
	if err != nil {
		logrus.Errorf("%s", err)
		os.Exit(1)
	}
}
