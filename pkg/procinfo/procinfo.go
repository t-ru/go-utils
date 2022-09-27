/*
MIT License

Copyright (c) 2022 Thomas Rudolph

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package procinfo

import (
	"encoding/csv"
	"os"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

type procInfo struct {
	processId            int
	processName          string
	processCmdline       string
	processIsGoRun       bool
	parentProcessId      int
	parentProcessName    string
	parentProcessCmdline string
	parentProcessIsGoRun bool
}

func isGoRun(procName string, procCmdline string) bool {

	switch os := runtime.GOOS; os {
	case "windows":
		if procName != "go.exe" {
			return false
		}
	default:
		if procName != "go" {
			return false
		}
	}

	// Split the Cmdline
	//	- 	Problem:
	//		Command Line can contain quotes
	//		e.g."C:\Program Files\Go\bin\go.exe" run .\main.go
	//		a normal split results in:
	//			- field 1: "C:\Program
	//			- field 2: Files\Go\bin\go.exe"
	//			- field 3: run
	//			- field 4: .\main.go
	//	-	Solution:
	//		csv.reader
	//		this results in:
	//			- field 1: "C:\Program Files\Go\bin\go.exe"
	//			- field 2: "run"
	//			- field 3: ".\main.go"
	reader := csv.NewReader(strings.NewReader(procCmdline))
	reader.Comma = ' ' // space
	fields, _ := reader.Read()

	// command "go run" has at least a length of 3 --> "go run <file(s)>"
	if len(fields) < 3 {
		return false
	}

	// command is "go <args>" ... but is it "go run <file(s)>" ?
	if fields[1] != "run" {
		return false
	}

	// command is "go run <file(s)>"
	return true
}

func New(pid int) *procInfo {

	var p *process.Process
	pi := &procInfo{}

	pi.processId = pid
	pi.parentProcessId = os.Getppid()

	p, _ = process.NewProcess(int32(pi.processId))
	pi.processName, _ = p.Name()
	pi.processCmdline, _ = p.Cmdline()
	pi.processIsGoRun = isGoRun(pi.processName, pi.processCmdline)

	p, _ = process.NewProcess(int32(pi.parentProcessId))
	pi.parentProcessName, _ = p.Name()
	pi.parentProcessCmdline, _ = p.Cmdline()
	pi.parentProcessIsGoRun = isGoRun(pi.parentProcessName, pi.parentProcessCmdline)

	return pi

}

func (pi *procInfo) ProcessId() int {
	return pi.processId
}

func (pi *procInfo) ProcessName() string {
	return pi.processName
}

func (pi *procInfo) ProcessCmdline() string {
	return pi.processCmdline
}

func (pi *procInfo) ProcessIsGoRun() bool {
	return pi.processIsGoRun
}

func (pi *procInfo) ParentProcessId() int {

	return pi.parentProcessId
}

func (pi *procInfo) ParentProcessName() string {
	return pi.parentProcessName
}

func (pi *procInfo) ParentProcessCmdline() string {
	return pi.parentProcessCmdline
}

func (pi *procInfo) ParentProcessIsGoRun() bool {
	return pi.parentProcessIsGoRun
}
