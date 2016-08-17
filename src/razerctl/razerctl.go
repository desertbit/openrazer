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
	"os"

	"github.com/urfave/cli"
)

const (
	Version = "1.0.0"
)

/*
func main() {
	cli.NewApp().Run(os.Args)

	// Initialize the client connection to the daemon.
	err := lib.Init()
	checkErrFatal(err)

	// Always close the daemon connection.
	defer lib.Close()

	devices, err := lib.GetDevices()
	checkErrFatal(err)

	for _, d := range devices {
		fmt.Printf("%+v\n", d)

		brightness, err := lib.GetBrightness(d.ID)
		checkErrFatal(err)
		fmt.Println("brightness:", brightness)
	}
}*/

func main() {
	app := cli.NewApp()
	app.Name = "razerctl"
	app.Version = Version
	app.Usage = "send query and control commands to the OpenRazer daemon"

	app.Commands = []cli.Command{
		{
			Name:    "devices",
			Aliases: []string{"l"},
			Usage:   "list all attached razer devices",
			Action:  getDevices,
		},
		{
			Name:    "device",
			Aliases: []string{"d"},
			Usage:   "show detailed information about a device",
			Action:  getDevice,
		},
	}

	app.Run(os.Args)
}

func getDevices(c *cli.Context) error {
	fmt.Println("added task: ", c.Args().First())
	return nil
}

func getDevice(c *cli.Context) error {
	fmt.Println("added task: ", c.Args().First())
	return nil
}
