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
	"C"

	"lib"
)

func main() {}

// It is the caller's responsibility to free the char array.
//export razer_get_last_error
func razer_get_last_error() *C.char {
	return C.CString(getLastErrString())
}

//export razer_init
func razer_init() C.int {
	err := lib.Init()
	if err != nil {
		setLastErr(err)
		return 0
	}

	return 1
}

//export razer_close
func razer_close() {
	lib.Close()
}

//export razer_get_brightness
func razer_get_brightness(id *C.char) (ret C.int, brightness C.int) {
	b, err := lib.GetBrightness(C.GoString(id))
	if err != nil {
		setLastErr(err)
		return 0, 0
	}

	return 1, C.int(b)
}
