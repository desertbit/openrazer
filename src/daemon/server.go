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

	"github.com/desertbit/pakt"
	"github.com/desertbit/pakt/tcp"

	log "github.com/Sirupsen/logrus"
)

func StartServer() error {
	// Create a new server.
	server, err := tcp.NewServer("127.0.0.1:42193")
	if err != nil {
		return err
	}

	// Set the handler function.
	server.OnNewSocket(onNewSocket)

	// Log.
	log.Infoln("daemon server listening...")

	// Start the server.
	server.Listen()

	return nil
}

func onNewSocket(s *pakt.Socket) {
	// Log as soon as the socket closed.
	s.OnClose(func(s *pakt.Socket) {
		log.Infof("client socket closed with remote address: %s", s.RemoteAddr())
	})

	// Set the call hook for logging purpose.
	s.SetCallHook(func(s *pakt.Socket, funcID string, c *pakt.Context) {
		log.WithFields(log.Fields{
			"remoteAddress": s.RemoteAddr(),
			"type":          funcID,
			"dataSize":      len(c.Data),
		}).Info("client request")
	})

	// Set the error hook for logging purpose.
	s.SetErrorHook(func(s *pakt.Socket, funcID string, err error) {
		log.WithFields(log.Fields{
			"remoteAddress": s.RemoteAddr(),
			"type":          funcID,
		}).Warningf("client request failed: %v", err)
	})

	// Register a remote callable function.
	// Optionally use s.RegisterFuncs to register multiple functions at once.
	s.RegisterFuncs(pakt.Funcs{
		"getDevices":    getDevices,
		"getDevice":     getDevice,
		"getBrightness": getBrightness,
		"getFnMode":     getFnMode,
		"getKeyRows":    getKeyRows,
		"getKeyColumns": getKeyColumns,
		"setBrightness": setBrightness,
		"setFnMode":     setFnMode,
	})

	// Signalize the socket that initialization is done.
	// Start accepting remote requests.
	s.Ready()

	// Log.
	log.Printf("new client socket with id: %s", s.ID())
}

func getDevices(c *pakt.Context) (interface{}, error) {
	devices := GetDevices()
	apiDevices := make(api.Devices, len(devices))

	for i, d := range devices {
		apiDevices[i] = d.ToApiDevice()
	}

	return apiDevices, nil
}

func getDevice(c *pakt.Context) (interface{}, error) {
	var id string
	err := c.Decode(&id)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(id)
	if err != nil {
		return nil, err
	}

	return d.ToApiDevice(), nil
}

func getBrightness(c *pakt.Context) (interface{}, error) {
	var id string
	err := c.Decode(&id)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(id)
	if err != nil {
		return nil, err
	}

	b, err := d.GetBrightness()
	if err != nil {
		return nil, err
	}

	return b, nil
}

func setBrightness(c *pakt.Context) (interface{}, error) {
	var v api.SetBrightnessValue
	err := c.Decode(&v)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(v.ID)
	if err != nil {
		return nil, err
	}

	err = d.SetBrightness(v.Brightness)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getFnMode(c *pakt.Context) (interface{}, error) {
	var id string
	err := c.Decode(&id)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(id)
	if err != nil {
		return nil, err
	}

	fn, err := d.GetFnMode()
	if err != nil {
		return nil, err
	}

	return fn, nil
}

func setFnMode(c *pakt.Context) (interface{}, error) {
	var v api.SetFnModeValue
	err := c.Decode(&v)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(v.ID)
	if err != nil {
		return nil, err
	}

	err = d.SetFnMode(v.FnMode)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func getKeyRows(c *pakt.Context) (interface{}, error) {
	var id string
	err := c.Decode(&id)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(id)
	if err != nil {
		return nil, err
	}

	r, err := d.GetKeyRows()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func getKeyColumns(c *pakt.Context) (interface{}, error) {
	var id string
	err := c.Decode(&id)
	if err != nil {
		return nil, err
	}

	d, err := GetDevice(id)
	if err != nil {
		return nil, err
	}

	co, err := d.GetKeyColumns()
	if err != nil {
		return nil, err
	}

	return co, nil
}
