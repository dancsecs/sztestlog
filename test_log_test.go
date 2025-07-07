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

package sztestlog_test

import (
	"testing"

	"github.com/dancsecs/szlog"
	"github.com/dancsecs/sztest"
	"github.com/dancsecs/sztestlog"
)

func TestSzTestLog_CaptureNothing(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureNothing(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))
}

func TestSzTestLog_CaptureStdout(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureStdout(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stdout()
}

func TestSzTestLog_CaptureLog(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureLog(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Log()
}

func TestSzTestLog_CaptureLogAndStdout(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureLogAndStdout(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stdout()
	tstChk.Log()
}

func TestSzTestLog_CaptureLogAndStderr(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureLogAndStderr(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Log()
	tstChk.Stderr()
}

func TestSzTestLog_CaptureLogAndStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureLogAndStderrAndStdout(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stdout()
	tstChk.Log()
	tstChk.Stderr()
}

func TestSzTestLog_CaptureLogWithStderr(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureLogWithStderr(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Log()
}

func TestSzTestLog_CaptureLogWithStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureLogWithStderrAndStdout(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stdout()
	tstChk.Log()
}

func TestSzTestLog_CaptureStderr(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureStderr(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stderr()
}

func TestSzTestLog_CaptureStderrAndStdout(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureStderrAndStdout(t, szlog.LevelAll)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stdout()
	tstChk.Stderr()
}

func TestSzTestLog_CaptureAll(t *testing.T) {
	chk := sztest.CaptureNothing(t)
	defer chk.Release()

	szlog.Reset()

	origLevel := szlog.SetLevel(szlog.LevelNone)
	defer func() {
		szlog.SetLevel(origLevel)
	}()

	tstChk := sztestlog.CaptureAll(t)
	defer tstChk.Release()

	chk.NotNil(tstChk)
	chk.Int(int(szlog.Level()), int(szlog.LevelAll))

	tstChk.Stdout()
	tstChk.Log()
	tstChk.Stderr()
}
