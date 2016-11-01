package callback

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Callback interface {
	// Timestamp of the received message
	Timestamp() time.Time
}

type ContentTypeError string

func (c ContentTypeError) Error() string {
	return string(c)
}

type MethodError string

func (c MethodError) Error() string {
	return string(c)
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
	Bidirectional  bool    `json:"ack"`
}

func parseCallback(r *http.Request, cb *callback) error {
	contentType := r.Header.Get("Content-Type")

	if r.Method == "POST" {
		if strings.Compare("application/json", contentType) == 0 {
			// unmarshall to callback struct
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(cb); err != nil {
				return err
			}

			return nil
		}

		return ContentTypeError(fmt.Sprintf("Unsupported content type: %s", contentType))
	} else {
		return MethodError(fmt.Sprintf("Unsupported request method: %s", r.Method))
	}

	return nil
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
