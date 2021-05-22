package watcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/sirupsen/logrus"

	"github.com/ArthurKnoep/vitemadose-notifier/internal/config"
)

type (
	watcher struct {
		logger *logrus.Logger
		cfg    *config.Config
	}

	center struct {
		name      string
		url       string
		address   string
		available int
	}
)

const chronodose = "chronodose"

func (w *watcher) fetchData() (*ViteMaDose, error) {
	resp, err := http.Get(fmt.Sprintf("https://vitemadose.gitlab.io/vitemadose/%d.json", w.cfg.DepartmentNumber))
	if err != nil {
		w.logger.WithError(err).Error("Unable to fetch data from vitemadose")
		return nil, err
	} else if resp.StatusCode >= 400 {
		w.logger.WithError(err).Error("Invalid status code returned from vitemadose")
		return nil, errors.New("invalid status code")
	}
	decoder := json.NewDecoder(resp.Body)
	var data ViteMaDose
	if err := decoder.Decode(&data); err != nil {
		w.logger.WithError(err).Error("Unable to parse data from vitemadose")
		return nil, err
	}
	return &data, err
}

func (w *watcher) startJob() {
	data, err := w.fetchData()
	if err != nil {
		return
	}

	var availableCenters []center
	for _, centerAvailable := range data.CentresDisponibles {
		var chronodoseCount = 0
		for _, appointment := range centerAvailable.AppointmentSchedules {
			if appointment.Name == chronodose && appointment.Total > 0 {
				chronodoseCount = appointment.Total
			}
		}
		if chronodoseCount > 0 {
			availableCenters = append(availableCenters, center{
				name:      centerAvailable.Nom,
				url:       centerAvailable.Url,
				address:   centerAvailable.Metadata.Address,
				available: chronodoseCount,
			})
		}
	}

	if len(availableCenters) == 0 {
		w.logger.Info("No chronodose available")
		return
	}
	_ = beeep.Notify("Vitemadose Notifier", fmt.Sprintf("%d centers have chronodose slot available", len(availableCenters)), "")
	w.logger.Infof("%d centers have chronodose slot available", len(availableCenters))
	for _, availableCenter := range availableCenters {
		w.logger.Infof("\t%s (%s): %d", availableCenter.name, availableCenter.url, availableCenter.available)
	}
}

func Watch(logger *logrus.Logger, cfg *config.Config) {
	w := &watcher{
		logger: logger,
		cfg:    cfg,
	}
	logger.Info("Start watching...")
	for {
		logger.Info("Starting data fetching")
		w.startJob()

		time.Sleep(cfg.ReloadTime)
	}
}
