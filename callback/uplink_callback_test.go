package callback

import (
	"net/http"
	"strings"
	"testing"
)

func TestParseUplinkCallback(t *testing.T) {
	// expected result from parsing JSON data
	ref := &UplinkCallback{
		callback: &callback{
			TimestampEpoch: 1477925692,
			DeviceID:       "01234567",
			IsDuplicate:    true,
			SNR:            1.23,
			RSSI:           3.21,
			AverageSNR:     4.56,
			StationID:      "76543210",
			Data:           "hello world",
			Latitude:       51,
			Longitude:      -1,
			SequenceNumber: 100,
		},
	}

	// build http callback request
	body := strings.NewReader(`{
		"time": 1477925692,
		"device": "01234567",
		"duplicate": true,
		"snr": 1.23,
		"rssi": 3.21,
		"avgSnr": 4.56,
		"station": "76543210",
		"data": "hello world",
		"lat": 51,
		"lng": -1,
		"seqNumber": 100
	}`)

	r, _ := http.NewRequest("POST", "/", body)
	r.Header.Set("ContentType", "application/json")

	// parse request body to UplinkCallback
	c, err := ParseUplinkCallback(r)
	if err != nil {
		t.Fatalf("Failed to parse uplink callback with: %v", err)
	}

	// test that parsed callback equals reference
	if !c.Equal(ref) {
		t.Fatalf("Parsed callback does not match expected value: %#v", c.callback)
	}
}
