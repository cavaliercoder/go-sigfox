package callback

import (
	"net/http"
	"time"
)

type UplinkCallback struct {
	callback *callback
}

func ParseUplinkCallback(r *http.Request) (*UplinkCallback, error) {
	c := &UplinkCallback{callback: &callback{}}
	if err := parseCallback(r, c.callback); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *UplinkCallback) Equal(u *UplinkCallback) bool {
	return c.callback.Equal(u.callback)
}

// Timestamp of the received message
func (c *UplinkCallback) Timestamp() time.Time {
	return time.Unix(c.callback.TimestampEpoch, 0)
}

// Device identifier in hexidecimal
func (c *UplinkCallback) DeviceID() string {
	return c.callback.DeviceID
}

// Whether message is a duplicate one, meanding that the backend has already
// processes this message from a different base station.
func (c *UplinkCallback) IsDuplicate() bool {
	return c.callback.IsDuplicate
}

// Signal to Noise Ration
func (c *UplinkCallback) SNR() float64 {
	return c.callback.SNR
}

// Average Signal to Noise Ratio for the last 25 messages
func (c *UplinkCallback) AverageSNR() float64 {
	return c.callback.AverageSNR
}

// Received Signal Strength Indicator in dB
func (c *UplinkCallback) RSSI() float64 {
	return c.callback.RSSI
}

// Base station identifier in hexidecimal
func (c *UplinkCallback) StationID() string {
	return c.callback.StationID
}

// User data as a byte slice
func (c *UplinkCallback) Data() []byte {
	return []byte(c.callback.Data)
}

// Latitude of the base station that received the message
func (c *UplinkCallback) Latitude() int64 {
	return c.callback.Latitude
}

// Longitude of the base station that received the message
func (c *UplinkCallback) Longitude() int64 {
	return c.callback.Longitude
}

// Sequence number of the message if available
func (c *UplinkCallback) SequenceID() int64 {
	return c.callback.SequenceNumber
}
