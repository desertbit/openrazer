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

/*
#include <stdlib.h>

struct razer_device {
	char *id;
	char *device_id;
	char *name;
	char *serial;
	char *firmware_version;
};
*/
import "C"

import (
	"unsafe"

	"lib"
)

func main() {}

// Deallocates the memory previously allocated by a call to calloc, malloc, or realloc.
//export razer_free
func razer_free(ptr unsafe.Pointer) {
	C.free(ptr)
}

//export razer_device_free
func razer_device_free(d *C.struct_razer_device) {
	C.free(unsafe.Pointer(d.id))
	C.free(unsafe.Pointer(d.device_id))
	C.free(unsafe.Pointer(d.name))
	C.free(unsafe.Pointer(d.serial))
	C.free(unsafe.Pointer(d.firmware_version))
	C.free(unsafe.Pointer(d))
}

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

// It is the caller's responsibility to free the device memory with razer_device_free.
// Retruns nil on error.
//export razer_get_device
func razer_get_device(id *C.char) (device *C.struct_razer_device) {
	d, err := lib.GetDevice(C.GoString(id))
	if err != nil {
		setLastErr(err)
		return nil
	}

	device = (*C.struct_razer_device)(C.malloc(C.size_t(unsafe.Sizeof(C.struct_razer_device{}))))
	device.id = C.CString(d.ID)
	device.device_id = C.CString(d.DeviceID)
	device.name = C.CString(d.Name)
	device.serial = C.CString(d.Serial)
	device.firmware_version = C.CString(d.FirmwareVersion)

	return device
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
