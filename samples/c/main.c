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

#include <stdlib.h>
#include <stdio.h>
#include <openrazer.h>


void print_last_error()
{
    char *err = razer_get_last_error();
    printf("error: %s\n", err);
    free(err);
}


int main(int argc, char const *argv[])
{
    int ret;

    // Initialize the openrazer module.
    ret = razer_init();
    if (!ret) {
        print_last_error();
        return ret;
    }

    // Brightness.
    struct razer_get_brightness_return b_ret = razer_get_brightness("950a78cfe4049c41742b5ab0e0b62a60dd584b1d");
    if (!b_ret.r0) {
        print_last_error();
        return ret;
    }
    printf("brightness: %d\n", b_ret.r1);


    razer_close();
    return 0;
}
