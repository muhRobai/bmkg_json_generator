package service

import (
	"encoding/xml"
	"github.com/bmkg_json_generator/constant"
	"github.com/bmkg_json_generator/model"
	"io/ioutil"
	"net/http"
)

type InitApi struct{}

type NewApi interface {
	RecentEarthQuakes() (*model.InfoEarthQuakes, error)
	List(limit int) (*model.InfoEarthQuakesList, error)
}

func NewApiService() NewApi {
	c := InitApi{}
	return &c
}

func (c *InitApi) RecentEarthQuakes() (*model.InfoEarthQuakes, error) {
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

func (c *InitApi) List(limit int) (*model.InfoEarthQuakesList, error) {
	size := constant.GetLimit(limit)

	url := constant.LimitMore
	if size < constant.MaxLimit {
		url = constant.LimitLess
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var p model.InfoEarthQuakesList
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(bodyBytes, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
