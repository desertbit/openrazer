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
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func stringToSHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return strings.TrimSpace(hex.EncodeToString(h.Sum(nil)))
}

func readFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(b)), nil
}

func readIntFromFile(path string) (int, error) {
	s, err := readFromFile(path)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return i, nil
}
