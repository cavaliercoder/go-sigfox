package callback

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Callback interface {
	// Timestamp of the received message
	Timestamp() time.Time
}

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
}

func parseCallback(r *http.Request, c *callback) error {
	contentType := r.Header.Get("ContentType")

	if "application/json" == contentType {
		// unmarshall to callback struct
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(c); err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("Unsupported content type: %s", contentType)
}

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
