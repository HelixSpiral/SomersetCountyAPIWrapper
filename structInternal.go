package somersetcountywrapper

import (
	"encoding/json"
	"strings"
)

type dispatchString string

type dispatchLogInternal struct {
	CallDate             dispatchString `json:"CallDate,omitempty"`
	CallNum              dispatchString `json:"CallNum,omitempty"`
	CallTime             dispatchString `json:"CallTime,omitempty"`
	ReasonText           dispatchString `json:"ReasonText,omitempty"`
	ActionCode           dispatchString `json:"ActionCode,omitempty"`
	ActionDesc           dispatchString `json:"ActionDesc,omitempty"`
	Jurisdiction         dispatchString `json:"Jurisdiction,omitempty"`
	StreetName           dispatchString `json:"StreetName,omitempty"`
	StreetSuf            dispatchString `json:"StreetSuf,omitempty"`
	Officer              dispatchString `json:"Officer1,omitempty"`
	LastName             dispatchString `json:"LastName,omitempty"`
	FirstName            dispatchString `json:"FirstName,omitempty"`
	Unit                 dispatchString `json:"Unit,omitempty"`
	UnitType             dispatchString `json:"UnitType,omitempty"`
	UnitDesc             dispatchString `json:"UnitDesc,omitempty"`
	Dispatched           dispatchString `json:"Dispatched,omitempty"`
	DispatchedSec        dispatchString `json:"DispatchedSec,omitempty"`
	Enroute              dispatchString `json:"Enroute,omitempty"`
	EnrouteSec           dispatchString `json:"EnrouteSec,omitempty"`
	Arrived              dispatchString `json:"Arrived,omitempty"`
	ArrivedSec           dispatchString `json:"ArrivedSec,omitempty"`
	Cleared              dispatchString `json:"Cleared,omitempty"`
	ClearedSec           dispatchString `json:"ClearedSec,omitempty"`
	ArrivedAtHospital    dispatchString `json:"ArrivedAtHospital,omitempty"`
	ArrivedAtHospitalSec dispatchString `json:"ArrivedAtHospitalSec,omitempty"`
	ClearedHospital      dispatchString `json:"ClearedHospital,omitempty"`
	ClearedHospitalSec   dispatchString `json:"ClearedHospitalSec,omitempty"`
	BackInService        dispatchString `json:"BackInService,omitempty"`
	BackInServiceSec     dispatchString `json:"BackInServiceSec,omitempty"`
	InQuarters           dispatchString `json:"InQuarters,omitempty"`
	InQuartersSec        dispatchString `json:"InQuartersSec,omitempty"`
	Year                 dispatchString `json:"Year,omitempty"`
	Make                 dispatchString `json:"Make,omitempty"`
	Style                dispatchString `json:"Style,omitempty"`
	Color                dispatchString `json:"Color1,omitempty"`
	Model                dispatchString `json:"Model,omitempty"`
	RegState             dispatchString `json:"RegState,omitempty"`
	RegNumber            dispatchString `json:"RegNumber,omitempty"`
	PlateType            dispatchString `json:"PlateType,omitempty"`
}

// For some reason the dispatch API has a ton of whitespaces
// This custom unmarshaler trims them all
func (d *dispatchString) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*d = dispatchString(strings.TrimSpace(v))

	return nil
}
