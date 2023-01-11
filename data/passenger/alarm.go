package passenger

import "encoding/json"

type Alarm struct {
	PfReal *PfReal `json:"pfReal,omitempty"`
}

type PfReal struct {
	PfStationReal []PfStationReal `json:"pfStationReal,omitempty"`
}

type PfStationReal struct {
	LineID   *string   `json:"lineId,omitempty"`
	LineName *string   `json:"lineName,omitempty"`
	Stations []Station `json:"stations,omitempty"`
}

type Station struct {
	AllQuatity           *json.Number     `json:"allQuatity"`
	PsgIn                *string          `json:"psgIn,omitempty"`
	StationName          *string          `json:"stationName,omitempty"`
	AlarmLevel           *json.Number     `json:"alarmLevel"`
	PsgTurn              *string          `json:"psgTurn,omitempty"`
	StationControlStatus *int64           `json:"stationControlStatus,omitempty"`
	TransferStation      *TransferStation `json:"transferStation,omitempty"`
	StationID            *string          `json:"stationId,omitempty"`
}

type TransferStation string

const (
	换乘站 TransferStation = "换乘站"
	普通站 TransferStation = "普通站"
)
