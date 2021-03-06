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

	"api"
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
		{
			Name:    "set-brightness",
			Aliases: []string{"sb"},
			Usage:   "set the brightness in percent [0-100]",
			Action:  setBrightness,
		},
		{
			Name:    "set-fn-mode",
			Aliases: []string{"sfn"},
			Usage:   "enable or disable the fn mode",
			Action:  setFnMode,
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
		table.Append([]string{"BRIGHTNESS", strconv.Itoa(brightness) + "%"})
	} else if err != nil && err != api.ErrNotSupported {
		return err
	}

	fnModeEnabled, err := lib.GetFnMode(d.ID)
	if err == nil {
		table.Append([]string{"FN MODE ENABLED", strconv.FormatBool(fnModeEnabled)})
	} else if err != nil && err != api.ErrNotSupported {
		return err
	}

	keyRows, err := lib.GetKeyRows(d.ID)
	if err == nil {
		table.Append([]string{"KEY ROWS", strconv.Itoa(keyRows)})
	} else if err != nil && err != api.ErrNotSupported {
		return err
	}

	keyColumns, err := lib.GetKeyColumns(d.ID)
	if err == nil {
		table.Append([]string{"KEY COLUMNS", strconv.Itoa(keyColumns)})
	} else if err != nil && err != api.ErrNotSupported {
		return err
	}

	// Print the formatted output.
	table.Render()

	return nil
}

func setBrightness(c *cli.Context) error {
	id := c.Args().First()
	if len(id) == 0 {
		return fmt.Errorf("missing ID: specify the device ID")
	}

	b := c.Args().Get(1)
	if len(b) == 0 {
		return fmt.Errorf("missing brightness value: specify the brightness value in percent [0-100]")
	}

	bi, err := strconv.Atoi(b)
	if err != nil {
		return err
	}

	err = lib.Init()
	if err != nil {
		return err
	}

	err = lib.SetBrightness(id, bi)
	if err != nil {
		return err
	}

	return nil
}

func setFnMode(c *cli.Context) error {
	id := c.Args().First()
	if len(id) == 0 {
		return fmt.Errorf("missing ID: specify the device ID")
	}

	active := c.Args().Get(1)
	if len(active) == 0 {
		return fmt.Errorf("missing active value: specify the fn mode active value [0=disabled 1=enabled]")
	}

	activeI, err := strconv.Atoi(active)
	if err != nil {
		return err
	}
	if activeI != 0 && activeI != 1 {
		return fmt.Errorf("invalid active fn mode value: 0=disabled 1=enabled")
	}

	err = lib.Init()
	if err != nil {
		return err
	}

	err = lib.SetFnMode(id, (activeI == 1))
	if err != nil {
		return err
	}

	return nil
}
