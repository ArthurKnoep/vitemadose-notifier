package watcher

import "time"

type ViteMaDose struct {
	Version            int           `json:"version"`
	LastUpdated        time.Time     `json:"last_updated"`
	LastScrap          []interface{} `json:"last_scrap"`
	CentresDisponibles []struct {
		Departement string `json:"departement"`
		Nom         string `json:"nom"`
		Url         string `json:"url"`
		Location    struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
			City      string  `json:"city"`
			Cp        string  `json:"cp"`
		} `json:"location"`
		Metadata struct {
			Address       string `json:"address"`
			BusinessHours *struct {
				Lundi    *string `json:"lundi"`
				Mardi    *string `json:"mardi"`
				Mercredi *string `json:"mercredi"`
				Jeudi    *string `json:"jeudi"`
				Vendredi *string `json:"vendredi"`
				Samedi   *string `json:"samedi"`
				Dimanche *string `json:"dimanche"`
			} `json:"business_hours"`
			PhoneNumber string `json:"phone_number"`
		} `json:"metadata"`
		ProchainRdv                string      `json:"prochain_rdv"`
		Plateforme                 string      `json:"plateforme"`
		Type                       string      `json:"type"`
		AppointmentCount           int         `json:"appointment_count"`
		InternalId                 string      `json:"internal_id"`
		VaccineType                []string    `json:"vaccine_type"`
		AppointmentByPhoneOnly     bool        `json:"appointment_by_phone_only"`
		Erreur                     interface{} `json:"erreur"`
		LastScanWithAvailabilities time.Time   `json:"last_scan_with_availabilities"`
		RequestCounts              struct {
			Slots    int `json:"slots"`
			Resource int `json:"resource,omitempty"`
			Booking  int `json:"booking,omitempty"`
			Cabinets int `json:"cabinets,omitempty"`
			Motives  int `json:"motives,omitempty"`
		} `json:"request_counts"`
		AppointmentSchedules []struct {
			Name  string    `json:"name"`
			From  time.Time `json:"from"`
			To    time.Time `json:"to"`
			Total int       `json:"total"`
		} `json:"appointment_schedules"`
		Gid string `json:"gid"`
	} `json:"centres_disponibles"`
	CentresIndisponibles []struct {
		Departement string `json:"departement"`
		Nom         string `json:"nom"`
		Url         string `json:"url"`
		Location    struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
			City      string  `json:"city"`
			Cp        string  `json:"cp"`
		} `json:"location"`
		Metadata struct {
			Address       string `json:"address"`
			PhoneNumber   string `json:"phone_number,omitempty"`
			BusinessHours *struct {
				Lundi    *string `json:"lundi"`
				Mardi    *string `json:"mardi"`
				Mercredi *string `json:"mercredi"`
				Jeudi    *string `json:"jeudi"`
				Vendredi *string `json:"vendredi"`
				Samedi   *string `json:"samedi"`
				Dimanche *string `json:"dimanche"`
			} `json:"business_hours"`
		} `json:"metadata"`
		ProchainRdv                interface{} `json:"prochain_rdv"`
		Plateforme                 string      `json:"plateforme"`
		Type                       string      `json:"type"`
		AppointmentCount           int         `json:"appointment_count"`
		InternalId                 string      `json:"internal_id"`
		VaccineType                []string    `json:"vaccine_type"`
		AppointmentByPhoneOnly     bool        `json:"appointment_by_phone_only"`
		Erreur                     interface{} `json:"erreur"`
		LastScanWithAvailabilities *time.Time  `json:"last_scan_with_availabilities"`
		RequestCounts              *struct {
			Booking  int `json:"booking,omitempty"`
			Motives  int `json:"motives,omitempty"`
			Slots    int `json:"slots"`
			Resource int `json:"resource,omitempty"`
			Cabinets int `json:"cabinets,omitempty"`
		} `json:"request_counts"`
		AppointmentSchedules []struct {
			Name  string    `json:"name"`
			From  time.Time `json:"from"`
			To    time.Time `json:"to"`
			Total int       `json:"total"`
		} `json:"appointment_schedules"`
		Gid string `json:"gid"`
	} `json:"centres_indisponibles"`
}
