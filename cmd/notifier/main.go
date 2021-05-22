package main

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ArthurKnoep/vitemadose-notifier/internal/config"
	"github.com/ArthurKnoep/vitemadose-notifier/internal/watcher"
)

func main() {
	app := kingpin.New(path.Base(os.Args[0]), "A simple tool to watch vitemadose appointment slot")
	app.DefaultEnvars()
	app.HelpFlag.Short('h')

	logger := logrus.New()

	c, err := config.New(app)
	if err != nil {
		logger.WithError(err).Error("invalid configuration")
		app.Usage(os.Args[1:])
		os.Exit(1)
	}
	watcher.Watch(logger, c)
}
