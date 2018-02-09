package sigfox

import (
	"encoding/hex"
	"net/http"
	"time"
)

// UplinkCallback is a callback message from the SIGFOX servers of the uplink
// type.
type UplinkCallback struct {
	callback *callback
}

type UplinkHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type UplinkHandlerFunc func(*UplinkCallback)

func (f UplinkHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	callbackHandlerFunc(func(cb *callback) {
		f(&UplinkCallback{callback: cb})
	}).ServeHTTP(w, r)
}

func (c *UplinkCallback) Equal(u *UplinkCallback) bool {
	return c.callback.Equal(u.callback)
}

// Timestamp of the received message
func (c *UplinkCallback) Timestamp() time.Time {
	return time.Unix(c.callback.TimestampEpoch, 0)
}

// Device identifier in hexadecimal
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

// Base station identifier in hexadecimal
func (c *UplinkCallback) StationID() string {
	return c.callback.StationID
}

// User data as a byte slice
func (c *UplinkCallback) Data() []byte {
	b, err := hex.DecodeString(c.callback.Data)
	if err != nil {
		dprintf("Error decoding SIGFOX payload: %v", err)
	}

	return b
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
