/*
   Szerszam logging library: szlog.
   Copyright (C) 2024  Leslie Dancsecs

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

/*
Package sztestlog provides convenience functions to return a sztest.Chk object
while setting the szlog package to the desired log level for testing.
*/
package sztestlog

import (
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

func setLogLevel(ll szlog.LogLevel) func() error {
	origLevel := szlog.SetLevel(ll)

	return func() error {
		szlog.SetLevel(origLevel)

		return nil
	}
}

// CaptureNothing returns a new *sztest.Chk reference and setting the logging
// level in a single call.
func CaptureNothing(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureNothing(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureStdout returns a new *sztest.Chk reference and setting the logging
// level in a single call.
func CaptureStdout(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLog returns a new *sztest.Chk reference and setting the logging
// level in a single call.
func CaptureLog(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureLog(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogAndStdout returns a new *sztest.Chk reference and setting the
// logging level in a single call.
func CaptureLogAndStdout(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureLogAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogAndStderr returns a new *sztest.Chk reference and setting the
// logging level in a single call.
func CaptureLogAndStderr(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureLogAndStderr(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogAndStderrAndStdout returns a new *sztest.Chk reference and
// setting the logging level in a single call.
func CaptureLogAndStderrAndStdout(
	t *testing.T, ll szlog.LogLevel,
) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureLogAndStderrAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogWithStderr returns a new *sztest.Chk reference and setting
// the logging level in a single call.
func CaptureLogWithStderr(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureLogWithStderr(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogWithStderrAndStdout returns a new *sztest.Chk reference and
// setting the logging level in a single call.
func CaptureLogWithStderrAndStdout(
	t *testing.T, ll szlog.LogLevel,
) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureLogWithStderrAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureStderr returns a new *sztest.Chk reference and setting the logging
// level in a single call.
func CaptureStderr(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureStderr(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureStderrAndStdout returns a new *sztest.Chk reference and setting
// the logging level in a single call.
func CaptureStderrAndStdout(t *testing.T, ll szlog.LogLevel) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(ll)

	chk := sztest.CaptureStderrAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureAll returns a new *sztest.Chk reference capturing log, stderr and
// stdout setting the logging level to LevelAll in a single call.
func CaptureAll(
	t *testing.T,
) *sztest.Chk {
	t.Helper()

	restoreFunc := setLogLevel(szlog.LevelAll)

	chk := sztest.CaptureLogAndStderrAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}
