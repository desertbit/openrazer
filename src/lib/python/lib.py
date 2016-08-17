#!/usr/bin/python
#
# OpenRazer
# Copyright (c) 2016 Roland Singer <roland.singer@desertbit.com>
# Copyright (c) 2016 Michael Hegel <michihegel@gmail.com>
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA 02111-1307 USA
#

import ctypes
import collections

####################
### Named Tuples ###
####################

device = collections.namedtuple('Device', ['id', 'device_id', 'name', 'serial', 'firmware_version'])



###############
### Private ###
###############

def _raiseExceptionIf(condition):
    if condition:
        raise ValueError(get_last_error())



##############
### Public ###
##############

# Init connects to the daemon and initializes the library.
# This method has to succeed before calling any other module methods.
def init():
    global lib
    lib = ctypes.CDLL("openrazer.so")
    lib.razer_init.restype = ctypes.c_int
    ret = lib.razer_init()
    _raiseExceptionIf(ret != 1)

def close():
    lib.close()

def get_last_error():
    lib.razer_get_last_error.restype = ctypes.c_void_p
    ptr = lib.razer_get_last_error()
    errStrB = ctypes.cast(ptr, ctypes.c_char_p).value
    lib.razer_free(ptr)
    return bytearray(errStrB).decode('utf8')

def get_device(id):
        class ret_struct(ctypes.Structure):
            _fields_ = [("id",                 ctypes.c_char_p),
                        ("device_id",          ctypes.c_char_p),
                        ("name",               ctypes.c_char_p),
                        ("serial",             ctypes.c_char_p),
                        ("firmware_version",   ctypes.c_char_p)]
        lib.razer_get_device.argtypes = [ctypes.c_char_p]
        lib.razer_get_device.restype = ctypes.POINTER(ret_struct)
        ret = lib.razer_get_device(id.encode('utf-8'))
        _raiseExceptionIf(not ret)
        d = device(id=bytearray(ret.contents.id).decode('utf8'),
                    device_id=bytearray(ret.contents.device_id).decode('utf8'),
                    name=bytearray(ret.contents.name).decode('utf8'),
                    serial=bytearray(ret.contents.serial).decode('utf8'),
                    firmware_version=bytearray(ret.contents.firmware_version).decode('utf8'))
        lib.razer_device_free(ret)
        return d

def get_brightness(id):
    class ret_struct(ctypes.Structure):
        _fields_ = [("r0", ctypes.c_int),
                    ("r1", ctypes.c_int)]
    lib.razer_get_brightness.argtypes = [ctypes.c_char_p]
    lib.razer_get_brightness.restype = ret_struct
    ret = lib.razer_get_brightness(id.encode('utf-8'))
    _raiseExceptionIf(ret.r0 != 1)
    return ret.r1
