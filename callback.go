package sigfox

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// encapsulates all possible sigfox callback types
// see: https://backend.sigfox.com/apidocs/callback
type callback struct {
	TimestampEpoch int64   `json:"time"`
	DeviceID       string  `json:"device"`
	IsDuplicate    bool    `json:"duplicate"`
	SNR            float64 `json:"snr"`
	RSSI           float64 `json:"rssi"`
	AverageSNR     float64 `json:"avgSnr"`
	StationID      string  `json:"station"`
	Data           string  `json:"data"`
	Latitude       int64   `json:"lat"`
	Longitude      int64   `json:"lng"`
	SequenceNumber int64   `json:"seqNumber"`
	Bidirectional  bool    `json:"ack"`
}

// parseCallback parses a http.Request and returns a base sigfox callback
func parseCallback(r *http.Request, cb *callback) (int, error) {
	defer r.Body.Close()

	contentType := r.Header.Get("Content-Type")

	switch r.Method {
	case "POST":
		switch contentType {
		case "application/json":
			// unmarshall JSON to callback struct
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(cb); err != nil {
				return http.StatusBadRequest, err
			}

		default:
			return http.StatusBadRequest, fmt.Errorf("Unsupported content type: %s", contentType)
		}

	default:
		return http.StatusMethodNotAllowed, fmt.Errorf("Unsupported request method: %s", r.Method)
	}

	return 0, nil
}

// Equal returns true if the values of two callback structs are equal
func (c *callback) Equal(b *callback) bool {
	return c.TimestampEpoch == b.TimestampEpoch &&
		c.DeviceID == b.DeviceID &&
		c.IsDuplicate == b.IsDuplicate &&
		c.SNR == b.SNR &&
		c.RSSI == b.RSSI &&
		c.AverageSNR == b.AverageSNR &&
		c.StationID == b.StationID &&
		c.Data == b.Data &&
		c.Latitude == b.Latitude &&
		c.Longitude == b.Longitude &&
		c.SequenceNumber == b.SequenceNumber
}
