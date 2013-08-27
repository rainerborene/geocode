package geocode

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type GeoIP struct {
	Areacode    string
	City        string
	CountryCode string `json:"country_code"`
	IP          string
	Latitude    float64
	Longitude   float64
	MetroCode   string `json:"metro_code"`
	RegionCode  string `json:"region_code"`
	RegionName  string `json:"region_name"`
	ZIPCode     string
}

func Geocode(s string) (*GeoIP, error) {
	location := &GeoIP{}
	ip := net.ParseIP(s)
	if ip == nil {
		return nil, errors.New("geocode: not a valid textual representation of an IP address")
	}
	url := fmt.Sprintf("http://freegeoip.net/json/%s", ip.String())
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func BatchGeocode(ips []string) []*GeoIP {
	ch := make(chan *GeoIP, len(ips))
	responses := []*GeoIP{}

	for _, ip := range ips {
		go func(ip string) {
			geo, err := Geocode(ip)
			if err != nil {
				log.Fatal(err)
			}
			ch <- geo
		}(ip)
	}

	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == len(ips) {
				return responses
			}
		}
	}

	return responses
}
