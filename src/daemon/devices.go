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
	"fmt"
	"io/ioutil"
	"regexp"
	"sync"

	log "github.com/Sirupsen/logrus"
)

//#################//
//### Variables ###//
//#################//

var (
	driverRegExp *regexp.Regexp

	devices      = make(map[string]*Device)
	devicesMutex sync.Mutex
)

//####################//
//### Devices Type ###//
//####################//

type Devices []*Device

//##############//
//### Public ###//
//##############//

// UpdateDevices updates the devices and removes unattached devices.
func UpdateDevices() error {
	e, err := exists(DriverPath)
	if err != nil {
		return err
	} else if !e {
		return fmt.Errorf("hid-razer driver modules is not loaded")
	}

	files, err := ioutil.ReadDir(DriverPath)
	if err != nil {
		return err
	}

	// Find all current attached devices.
	var deviceIDs []string
	for _, d := range files {
		dirName := d.Name()
		if driverRegExp.MatchString(dirName) {
			deviceIDs = append(deviceIDs, dirName)
		}
	}

	devicesMutex.Lock()
	defer devicesMutex.Unlock()

	// Add new devices which aren't present already.
	for _, deviceID := range deviceIDs {
		d := NewDevice(deviceID)
		err = d.loadInfoFromDriver()
		if err != nil {
			log.Errorf("failed to load information from the driver: %v", err)
			continue
		}

		// Create a unique ID from the device type and the device serial.
		id := stringToSHA1(d.deviceType + d.serial)

		// Skip if device is already in the map.
		if _, ok := devices[id]; ok {
			continue
		}

		devices[id] = d
	}

	// Remove devices which aren't present anymore.
DeleteLoop:
	for id, d := range devices {
		for _, deviceID := range deviceIDs {
			if d.deviceID == deviceID {
				continue DeleteLoop
			}
		}

		delete(devices, id)
	}

	return nil
}

// GetDevices returns a list of all attached devices.
func GetDevices() Devices {
	devicesMutex.Lock()
	defer devicesMutex.Unlock()

	devicesSlice := make(Devices, len(devices))

	i := 0
	for _, device := range devices {
		devicesSlice[i] = device
		i++
	}

	return devicesSlice
}

// GetDevice returns the device specified by the ID.
// Returns nil if not found.
func GetDevice(id string) *Device {
	devicesMutex.Lock()
	defer devicesMutex.Unlock()

	return devices[id]
}

//###############//
//### Private ###//
//###############//

func init() {
	var err error

	driverRegExp, err = regexp.Compile(`.*:.*:.*\..*`)
	if err != nil {
		log.Fatalf("failed to compile the RegExp: %v", err)
	}
}
