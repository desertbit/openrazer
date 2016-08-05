package api

type Devices []Device

type Device struct {
	ID              string
	Name            string
	Serial          string
	FirmwareVersion string
}
