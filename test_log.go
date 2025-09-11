/*
   Szerszam logging library: szlog.
   Copyright (C) 2024-2025  Leslie Dancsecs

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

package sztestlog

import (
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
)

func setLevels(t *testing.T, args []string) func() error {
	t.Helper()

	origLevel := szlog.Level()
	origVerbose := szlog.Verbose()
	origLongLabels := szlog.LongLabels()
	origLanguage := szlog.Language()

	if len(args) == 0 {
		args = []string{
			"-vvvvvv",
			"--log",
			"All",
		}
	}

	remainingArgs, err := szlog.AbsorbArgs(args)

	if err != nil {
		t.Log("Setting args caused: ", err)
	} else if len(remainingArgs) != 0 {
		t.Log("Setting args unknown arg(s): ", remainingArgs)
	}

	return func() error {
		szlog.SetLevel(origLevel)
		szlog.SetVerbose(origVerbose)
		szlog.SetLongLabels(origLongLabels)
		_ = szlog.SetLanguage(origLanguage)

		return nil
	}
}

// CaptureNothing returns a new *sztest.Chk reference while setting the szlog
// logging level in a single call.  If no szlogArgs are provided then szlog
// defaults to szlog.LevelAll and maximum verbosity.
func CaptureNothing(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureNothing(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureStdout returns a new *sztest.Chk reference while setting the szlog
// logging level in a single call.  If no szlogArgs are provided then szlog
// defaults to szlog.LevelAll and maximum verbosity.
func CaptureStdout(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLog returns a new *sztest.Chk reference while setting the szlog
// logging level in a single call.  If no szlogArgs are provided then szlog
// defaults to szlog.LevelAll and maximum verbosity.
func CaptureLog(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureLog(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogAndStdout returns a new *sztest.Chk reference while setting the
// szlog logging level in a single call.  If no szlogArgs are provided then
// szlog defaults to szlog.LevelAll and maximum verbosity.
func CaptureLogAndStdout(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureLogAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogAndStderr returns a new *sztest.Chk reference while setting the
// szlog logging level in a single call.  If no szlogArgs are provided then
// szlog defaults to szlog.LevelAll and maximum verbosity.
func CaptureLogAndStderr(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureLogAndStderr(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogAndStderrAndStdout returns a new *sztest.Chk reference while
// setting the szlog logging level in a single call.  If no szlogArgs are
// provided then szlog defaults to szlog.LevelAll and maximum verbosity.
func CaptureLogAndStderrAndStdout(
	t *testing.T, szlogArgs ...string,
) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureLogAndStderrAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogWithStderr returns a new *sztest.Chk reference while setting the
// szlog logging level in a single call.  If no szlogArgs are provided then
// szlog defaults to szlog.LevelAll and maximum verbosity.
func CaptureLogWithStderr(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureLogWithStderr(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureLogWithStderrAndStdout returns a new *sztest.Chk reference while
// setting the szlog logging level in a single call.  If no szlogArgs are
// provided then szlog defaults to szlog.LevelAll and maximum verbosity.
func CaptureLogWithStderrAndStdout(
	t *testing.T, szlogArgs ...string,
) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureLogWithStderrAndStdout(t)

	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureStderr returns a new *sztest.Chk reference while setting the szlog
// logging level in a single call.  If no szlogArgs are provided then szlog
// defaults to szlog.LevelAll and maximum verbosity.
func CaptureStderr(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureStderr(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}

// CaptureStderrAndStdout returns a new *sztest.Chk reference while setting
// the szlog logging level in a single call.  If no szlogArgs are provided
// then szlog defaults to szlog.LevelAll and maximum verbosity.
func CaptureStderrAndStdout(t *testing.T, szlogArgs ...string) *sztest.Chk {
	t.Helper()

	restoreFunc := setLevels(t, szlogArgs)

	chk := sztest.CaptureStderrAndStdout(t)
	chk.PushPostReleaseFunc(restoreFunc)

	return chk
}
