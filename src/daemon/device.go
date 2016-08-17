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

func NewDevice(deviceID string) *Device {
	return &Device{
		deviceID: deviceID,
	}
}

func (d *Device) ID() string {
	return d.id
}

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
	d.id = stringToSHA1(deviceID + d.serial)

	return nil
}

func (d *Device) devicePath() string {
	return DriverPath + "/" + d.deviceID + "/"
}
