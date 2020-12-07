package model

import (
	"encoding/xml"
)

type InfoEarthQuakes struct {
	xmlName     xml.Name
	EarthQuakes EarthQuakes `xml:"gempa" json:"gampa"`
}

type EarthQuakes struct {
	Date        string     `xml:"Tanggal" json:"tanggal"`
	Time        string     `xml:"Jam" json:"jam"`
	Coordinates Coordinate `xml:"point" json:"point"`
	Lintang     string     `xml:"Lintang" json:"lintang"`
	Bujur       string     `xml:"Bujur" json:"bujur"`
	Magnitude   string     `xml:"Magnitude" json:"magnitude"`
	Symbol      string     `xml:"_symbol" json:"symbol"`
	Kedalaman   string     `xml:"Kedalaman" json:"kedalaman"`
	Wilayah     string     `xml:"Wilayah" json:"wilayah"`
	Wilayah1    string     `xml:"Wilayah1" json:"wilayah_1"`
	Wilayah2    string     `xml:"Wilayah2" json:"wilayah_2"`
	Wilayah3    string     `xml:"Wilayah3" json:"wilayah_3"`
	Wilayah4    string     `xml:"Wilayah4" json:"wilayah_4"`
	Wilayah5    string     `xml:"Wilayah5" json:"wilayah_5"`
	Potensi     string     `xml:"Potensi" json:"potensi"`
	Keterangan  string     `xml:"Keterangan" json:"keterangan"`
	Dirasakan   string     `xml:"Dirasakan" json:"dirasakan"`
	Posisi      string     `xml:"Posisi" json:"posisi"`
}

type Coordinate struct {
	Coordinates string `xml:"coordinates" json:"coordinate"`
}

type InfoEarthQuakesList struct {
	xmlName xml.Name
	List    []*EarthQuakes `xml:"gempa" json:"genmpa"`
}
