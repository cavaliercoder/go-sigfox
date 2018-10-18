package sigfox

import (
	"encoding/hex"
	"net/http"
	"time"
)

// UplinkCallback is a callback message from
// the SIGFOX servers of the uplink type.
type UplinkCallback struct {
	callback *callback
}

// UplinkHandler responds to Uplink HTTP request.
type UplinkHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// UplinkHandlerFunc is an adapter to allow the use of ordinary functions as Uplink HTTP handlers.
type UplinkHandlerFunc func(*UplinkCallback)

// ServeHTTP serve HTTP response.
func (f UplinkHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	callbackHandlerFunc(func(cb *callback) {
		f(&UplinkCallback{callback: cb})
	}).ServeHTTP(w, r)
}

// Equal returns true if the values of two callback structs are equal.
func (c *UplinkCallback) Equal(u *UplinkCallback) bool {
	return c.callback.Equal(u.callback)
}

// Timestamp returns timestamp of the received message.
func (c *UplinkCallback) Timestamp() time.Time {
	return time.Unix(c.callback.TimestampEpoch, 0)
}

// DeviceID returns hexadecimal device identifier.
func (c *UplinkCallback) DeviceID() string {
	return c.callback.DeviceID
}

// IsDuplicate reports whether message is a duplicate.
// Meanding that the backend has already processes
// message from a different base station.
func (c *UplinkCallback) IsDuplicate() bool {
	return c.callback.IsDuplicate
}

// SNR returns Signal to Noise Ration.
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
