package callback

import (
	"net/http"
	"net/http/httptest"
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

	// start hookup test server
	ts := httptest.NewServer(HTTPHandlerFunc(func(cb Callback) error {
		if uplink, ok := cb.(*UplinkCallback); !ok {
			t.Errorf("Expected UplinkCallback type but got %t", cb)
		} else if uplink == nil {
			t.Errorf("Expected UplinkCallback type but got nil")
		} else {
			if !uplink.callback.Equal(ref.callback) {
				t.Errorf("Parsed callback does not match expected value: %#v", uplink.callback)
			}
		}

		return nil
	}))

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

	// parse request body to UplinkCallback
	if resp, err := http.DefaultClient.Post(ts.URL, "application/json", body); err != nil {
		t.Fatalf("Error: %v", err)
	} else if resp.StatusCode < 200 || resp.StatusCode > 299 {
		t.Fatalf("Received non 200 response")
	}
}
