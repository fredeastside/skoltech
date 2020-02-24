package devices

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

//Device structure
type Device struct {
	UID      string            `json:"uid"`
	ID       string            `json:"ap_id"`
	Version  string            `json:"version"`
	Requests []*deviceRequests `json:"probe_requests"`
}

type deviceRequests struct {
	MAC       string `json:"mac"`
	Timestamp string `json:"timestamp,omitempty"`
	BSSID     string `json:"bssid,omitempty"`
	SSID      string `json:"ssid,omitempty"`
}

//NewDevice constructor for device struct
func NewDevice() *Device {
	uid := uuid.NewV4()

	return &Device{
		UID: uid.String(),
	}
}

//Validate device
func (d *Device) Validate() error {
	if d.ID == "" {
		return errors.New("device has empty app id")
	}

	if len(d.Requests) == 0 {
		return errors.New("device has empty probe requests")
	}

	return nil
}

//Transform prepares data for resending
func (d *Device) Transform() error {
	for _, r := range d.Requests {
		if r.MAC == "" {
			return errors.New("device has empty MAC")
		}

		if r.BSSID == "" {
			r.BSSID = "FF-FF-FF-FF-FF-FF"
		}

		if r.SSID == "" {
			r.SSID = "Unknown"
		}
	}

	return nil
}
