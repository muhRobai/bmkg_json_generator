package service

import (
	"bmkg/bmkg_api/model"
	"context"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type InitApi struct{}

type NewApi interface {
	RecentEarthQuakes(context.Context) (*model.InfoEarthQuakes, error)
}

func NewApiService() NewApi {
	c := InitApi{}
	return &c
}

func (c *InitApi) RecentEarthQuakes(ctx context.Context) (*model.InfoEarthQuakes, error) {
	resp, err := http.Get("https://data.bmkg.go.id/eqmap.gif")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var p model.InfoEarthQuakes
	err = xml.Unmarshal(bodyBytes, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
