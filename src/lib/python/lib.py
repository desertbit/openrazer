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


def init():
    global _lib
    _lib = ctypes.CDLL("../c/openrazer.so")
    razer_init            = _lib.razer_init
    razer_init.restype    = ctypes.c_bool
    return razer_init()

def close():
    _lib.close()

def get_last_error():
    razer_get_last_error            = _lib.razer_get_last_error
    razer_get_last_error.restype    = ctypes.c_char_p
    return razer_get_last_error()
