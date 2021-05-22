package config

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

type Config struct {
	DepartmentNumber int
	ReloadTime       time.Duration
}

func (c *Config) String() string {
	res, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(res)
}

func New(app *kingpin.Application) (*Config, error) {
	var c Config

	app.Flag("department", "The department number to watch").
		Short('d').
		Required().
		IntVar(&c.DepartmentNumber)
	app.Flag("reload", "The time to wait before reloading the data").
		Short('r').
		Default("5m").
		DurationVar(&c.ReloadTime)

	_, err := app.Parse(os.Args[1:])

	if c.DepartmentNumber <= 0 || c.DepartmentNumber > 95 {
		return nil, errors.New("invalid department number")
	}
	return &c, err
}
