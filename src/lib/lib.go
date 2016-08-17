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

package lib

import (
	"fmt"

	"github.com/desertbit/pakt"
	"github.com/desertbit/pakt/tcp"

	"api"
)

var (
	socket *pakt.Socket
)

// Init creates a new client socket.
func Init() error {
	var err error

	// Create a new client.
	socket, err = tcp.NewClient("127.0.0.1:42193")
	if err != nil {
		return fmt.Errorf("failed to connect to daemon: %v", err)
	}

	// Signalize the socket that initialization is done.
	// Start accepting remote requests.
	socket.Ready()

	return nil
}

// Close closes the socket connection.
func Close() {
	if socket == nil {
		return
	}

	socket.Close()
}

// GetDevices returns a slice of current devices.
func GetDevices() (api.Devices, error) {
	// Call the server function in order to get an array of devices.
	c, err := socket.Call("getDevices")
	if err != nil {
		return nil, err
	}

	// Decode the return value.
	var devices api.Devices
	err = c.Decode(&devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// GetDevice returns one device that's specified by an id.
func GetDevice(id string) (*api.Device, error) {
	// Call the server function in order to get a specific device.
	c, err := socket.Call("getDevice", id)
	if err != nil {
		return nil, err
	}

	// Decode the return value.
	var device api.Device
	err = c.Decode(&device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

// GetBrightness obtains the brightness value in percent (0%-100%).
func GetBrightness(id string) (int, error) {
	c, err := socket.Call("getBrightness", id)
	if err != nil {
		return 0, err
	}

	// Decode the return value.
	var b int
	err = c.Decode(&b)
	if err != nil {
		return 0, err
	}

	return b, nil
}

// GetFnMode obtains the value of fn_mode.
func GetFnMode(id string) (bool, error) {
	c, err := socket.Call("getFnMode", id)
	if err != nil {
		return false, err
	}

	// Decode the return value.
	var fn bool
	err = c.Decode(&fn)
	if err != nil {
		return false, err
	}

	return fn, nil
}

// GetKeyRows obtains the (internal) amount of rows on the keyboard.
func GetKeyRows(id string) (int, error) {
	c, err := socket.Call("getKeyRows", id)
	if err != nil {
		return 0, err
	}

	// Decode the return value.
	var r int
	err = c.Decode(&r)
	if err != nil {
		return 0, err
	}

	return r, nil
}

// GetKeyColumns obtains the (internal) amount of columns on the keyboard.
func GetKeyColumns(id string) (int, error) {
	c, err := socket.Call("getKeyColumns", id)
	if err != nil {
		return 0, err
	}

	// Decode the return value.
	var co int
	err = c.Decode(&co)
	if err != nil {
		return 0, err
	}

	return co, nil
}
