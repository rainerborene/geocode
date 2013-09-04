package geocode

import "fmt"
import "testing"

func TestGeocode(t *testing.T) {
	loc, err := Geocode("189.59.162.230")
	if err != nil {
		t.Errorf("request failed: %s", err)
	}

	str := fmt.Sprintf("%s %s %s", loc.CountryCode, loc.RegionName, loc.City)

	if loc.String() != str {
		t.Errorf("wrong string representation: %s", loc)
	}
}
