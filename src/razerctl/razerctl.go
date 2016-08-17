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
	"strconv"

	"lib"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

const (
	Version = "1.0.0"
)

func main() {
	// Always close the daemon connection.
	defer lib.Close()

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
	err := lib.Init()
	if err != nil {
		return err
	}

	devices, err := lib.GetDevices()
	if err != nil {
		return err
	}

	// Format the output.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "NAME", "DEVICE ID"})

	for _, d := range devices {
		table.Append([]string{d.ID, d.Name, d.DeviceID})
	}

	// Print the formatted output.
	table.Render()

	return nil
}

func getDevice(c *cli.Context) error {
	id := c.Args().First()
	if len(id) == 0 {
		return fmt.Errorf("missing ID: specify the device ID")
	}

	err := lib.Init()
	if err != nil {
		return err
	}

	d, err := lib.GetDevice(id)
	if err != nil {
		return err
	}

	// Format the output.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")

	table.Append([]string{"NAME", d.Name})
	table.Append([]string{"ID", d.ID})
	table.Append([]string{"DEVICE ID", d.DeviceID})
	table.Append([]string{"FIRMWARE VERSION", d.FirmwareVersion})
	table.Append([]string{"SERIAL", d.Serial})

	brightness, err := lib.GetBrightness(d.ID)
	if err == nil {
		table.Append([]string{"Brightness", strconv.Itoa(brightness) + "%"})
	} else if err != nil {
		// TODO
	}

	// Print the formatted output.
	table.Render()

	return nil
}
