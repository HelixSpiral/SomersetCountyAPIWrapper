package somersetcountywrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const URL = "https://api.somersetcounty-me.org/DispatchLog/"

func NewWrapper() *SomersetWrapper {
	return &SomersetWrapper{}
}

func (d *SomersetWrapper) GetDispatch(date string) ([]DispatchLog, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", URL, date), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ACCEPT", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// This definitely seems weird, but we unmarshal it once with custom strings definitions so we can trim whitespace.
	// Then below we remarshal it, and unmarshal it again to the struct we're returning externally.
	var dispatchListInternal []dispatchLogInternal
	err = json.Unmarshal(body, &dispatchListInternal)
	if err != nil {
		return nil, err
	}

	jsonStripped, err := json.Marshal(dispatchListInternal)
	if err != nil {
		return nil, err
	}

	var dispatchList []DispatchLog
	err = json.Unmarshal(jsonStripped, &dispatchList)
	if err != nil {
		return nil, err
	}

	return dispatchList, nil
}
