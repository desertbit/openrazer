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
	"github.com/desertbit/pakt"
	"github.com/desertbit/pakt/tcp"

	"api"
)

var (
	socket *pakt.Socket
)

// Initialize creates a new client socket.
func Initialize() error {
	var err error

	// Create a new client.
	socket, err = tcp.NewClient("127.0.0.1:42193")
	if err != nil {
		return err
	}

	// Set a function which is triggered as soon as the socket closed.
	// Optionally use the s.ClosedChan channel.
	socket.OnClose(func(s *pakt.Socket) {
		Log.Errorf("Daemon connection lost.")
	})

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
