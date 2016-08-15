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
	"github.com/desertbit/pakt"
	"github.com/desertbit/pakt/tcp"

	log "github.com/Sirupsen/logrus"
)

func main() {

	err := startServer()
	if err != nil {
		log.Fatalf("Daemon server failed: %v\n", err)
	}
}

func startServer() error {
	// Create a new server.
	server, err := tcp.NewServer("127.0.0.1:42193")
	if err != nil {
		return err
	}

	// Set the handler function.
	server.OnNewSocket(onNewSocket)

	// Log.
	log.Infoln("Daemon server listening...")

	// Start the server.
	server.Listen()

	return nil
}

func onNewSocket(s *pakt.Socket) {
	// Log as soon as the socket closed.
	s.OnClose(func(s *pakt.Socket) {
		log.Errorf("client socket closed with id: %s", s.ID())
	})

	// Register a remote callable function.
	// Optionally use s.RegisterFuncs to register multiple functions at once.
	s.RegisterFuncs(pakt.Funcs{
		"getDevices": getDevices,
		"getDevice":  getDevice,
	})

	// Signalize the socket that initialization is done.
	// Start accepting remote requests.
	s.Ready()

	// Log.
	log.Printf("new client socket with id: %s", s.ID())
}

func getDevices(c *pakt.Context) (interface{}, error) {

	return nil, nil
}

func getDevice(c *pakt.Context) (interface{}, error) {
	return nil, nil
}
