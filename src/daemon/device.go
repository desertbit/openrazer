/*
 * OpenRazer
 * Copyright (c) 2016 Roland Singer <roland.singer@desertbit.com>
 * Copyright (c) 2016 Michael Hegel <michihegel@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307 USA
 */

package main

import (
	"api"
	"fmt"
	"strconv"
	"strings"
)

//####################//
//### Devices Type ###//
//####################//

type Device struct {
	id              string
	deviceID        string
	name            string
	serial          string
	firmwareVersion string
}

// NewDevice returns a new Device with the given deviceID filled in.
func NewDevice(deviceID string) *Device {
	return &Device{
		deviceID: deviceID,
	}
}

// ID returns the ID of the given Device.
func (d *Device) ID() string {
	return d.id
}

// ToApiDevice transforms a device to a device struct, that's defined in api.go.
func (d *Device) ToApiDevice() *api.Device {
	return &api.Device{
		ID:              d.id,
		DeviceID:        d.deviceID,
		Name:            d.name,
		Serial:          d.serial,
		FirmwareVersion: d.firmwareVersion,
	}
}

// GetBrightness returns the brightness faktor in percent (0%-100%).
func (d *Device) GetBrightness() (int, error) {
	b, err := readIntFromFile(d.devicePath() + "brightness")
	if err != nil {
		return 0, err
	}

	if b < 0 || b > 255 {
		return 0, fmt.Errorf("invalid brightness value: %v", b)
	}

	// Transform to percent.
	b = int(float64(b) / 2.55)

	return b, nil
}

// SetBrightness writes the brightness faktor in percent (0%-100%).
func (d *Device) SetBrightness(b int) error {
	if b < 0 || b > 100 {
		return fmt.Errorf("invalid brightness value: %v", b)
	}

	// Transform from percent and convert to string.
	v := strconv.Itoa(int(float64(b) * 2.55))

	err := writeToFile(d.devicePath()+"brightness", v)
	if err != nil {
		return err
	}

	return nil
}

// GetFnMode returns the current value of the fn_mode, it's either 0 or 1.
func (d *Device) GetFnMode() (bool, error) {
	fn, err := readIntFromFile(d.devicePath() + "fn_mode")
	if err != nil {
		return false, err
	}

	return (fn == 1), nil
}

// SetFnMode writes the  faktor in percent (0%-100%).
func (d *Device) SetFnMode(a bool) error {
	i := 0
	if a {
		i = 1
	}

	v := strconv.Itoa(i)

	err := writeToFile(d.devicePath()+"fn_mode", v)
	if err != nil {
		return err
	}

	return nil
}

// GetKeyRows returns the (internal) amount of rows on the keyboard.
func (d *Device) GetKeyRows() (int, error) {
	r, err := readIntFromFile(d.devicePath() + "get_key_rows")
	if err != nil {
		return 0, err
	}

	return r, nil
}

// GetKeyColumns returns the (internal) amount of columns on the keyboard.
func (d *Device) GetKeyColumns() (int, error) {
	co, err := readIntFromFile(d.devicePath() + "get_key_columns")
	if err != nil {
		return 0, err
	}

	return co, nil
}

//###############//
//### Private ###//
//###############//

func (d *Device) init() error {
	var err error
	devicePath := d.devicePath()

	d.name, err = readFromFile(devicePath + "device_type")
	if err != nil {
		return err
	}

	d.serial, err = readFromFile(devicePath + "get_serial")
	if err != nil {
		return err
	}

	d.firmwareVersion, err = readFromFile(devicePath + "get_firmware_version")
	if err != nil {
		return err
	}

	// Extract the real device ID.
	var deviceID string
	pos := strings.LastIndex(d.deviceID, ":")
	if pos >= 0 {
		pos2 := strings.LastIndex(d.deviceID, ".")
		if pos2 >= 0 {
			deviceID = d.deviceID[pos+1 : pos2]
		}
	}

	// Create a unique ID from the device ID and the device serial.
	d.id = stringToCRC64(deviceID + d.serial)

	return nil
}

func (d *Device) devicePath() string {
	return DriverPath + "/" + d.deviceID + "/"
}
