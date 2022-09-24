package somersetcountywrapper

type SomersetWrapper struct{}

type DispatchLog struct {
	CallDate             string `json:"CallDate,omitempty"`
	CallNum              string `json:"CallNum,omitempty"`
	CallTime             string `json:"CallTime,omitempty"`
	ReasonText           string `json:"ReasonText,omitempty"`
	ActionCode           string `json:"ActionCode,omitempty"`
	ActionDesc           string `json:"ActionDesc,omitempty"`
	Jurisdiction         string `json:"Jurisdiction,omitempty"`
	StreetName           string `json:"StreetName,omitempty"`
	StreetSuf            string `json:"StreetSuf,omitempty"`
	Officer              string `json:"Officer1,omitempty"`
	LastName             string `json:"LastName,omitempty"`
	FirstName            string `json:"FirstName,omitempty"`
	Unit                 string `json:"Unit,omitempty"`
	UnitType             string `json:"UnitType,omitempty"`
	UnitDesc             string `json:"UnitDesc,omitempty"`
	Dispatched           string `json:"Dispatched,omitempty"`
	DispatchedSec        string `json:"DispatchedSec,omitempty"`
	Enroute              string `json:"Enroute,omitempty"`
	EnrouteSec           string `json:"EnrouteSec,omitempty"`
	Arrived              string `json:"Arrived,omitempty"`
	ArrivedSec           string `json:"ArrivedSec,omitempty"`
	Cleared              string `json:"Cleared,omitempty"`
	ClearedSec           string `json:"ClearedSec,omitempty"`
	ArrivedAtHospital    string `json:"ArrivedAtHospital,omitempty"`
	ArrivedAtHospitalSec string `json:"ArrivedAtHospitalSec,omitempty"`
	ClearedHospital      string `json:"ClearedHospital,omitempty"`
	ClearedHospitalSec   string `json:"ClearedHospitalSec,omitempty"`
	BackInService        string `json:"BackInService,omitempty"`
	BackInServiceSec     string `json:"BackInServiceSec,omitempty"`
	InQuarters           string `json:"InQuarters,omitempty"`
	InQuartersSec        string `json:"InQuartersSec,omitempty"`
	Year                 string `json:"Year,omitempty"`
	Make                 string `json:"Make,omitempty"`
	Style                string `json:"Style,omitempty"`
	Color                string `json:"Color1,omitempty"`
	Model                string `json:"Model,omitempty"`
	RegState             string `json:"RegState,omitempty"`
	RegNumber            string `json:"RegNumber,omitempty"`
	PlateType            string `json:"PlateType,omitempty"`
}
