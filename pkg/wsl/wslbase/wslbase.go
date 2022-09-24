/*
MIT License

# Copyright (c) 2022 Thomas Rudolph

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
package wslbase

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/t-ru/go-utils/pkg/linebreakconverter"
)

func WslExecute(stdoutSilent bool, stderrSilent bool, args ...string) (stdout string, stderr string, retcode int, err error) {

	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer
	var cmd *exec.Cmd
	var cmdResult error
	var cmdStdout string
	var cmdStderr string
	var cmdRetcode int

	cmd = exec.Command("wsl.exe", args...)

	cmd.Stdin = os.Stdin

	if stdoutSilent {
		cmd.Stdout = &stdoutBuffer
	} else {
		// os.Stdout = crlf, stdoutBuffer = lf
		cmd.Stdout = io.MultiWriter(linebreakconverter.ConvertLinebreakToWindowsWriter(os.Stdout), &stdoutBuffer)
		// os.Stdout = crlf, stdoutBuffer = crlf
		//cmd.Stdout = linebreakconverter.ConvertLinebreakToWindowsWriter(io.MultiWriter(os.Stdout, &stdoutBuffer))
	}

	if stderrSilent {
		cmd.Stderr = &stderrBuffer
	} else {
		// os.Stdout = crlf, stderrBuffer = lf
		cmd.Stderr = io.MultiWriter(linebreakconverter.ConvertLinebreakToWindowsWriter(os.Stderr), &stderrBuffer)
		// os.Stderr = crlf, stderrBuffer = crlf
		// cmd.Stderr = linebreakconverter.ConvertLinebreakToWindowsWriter(io.MultiWriter(os.Stderr, &stderrBuffer))
	}

	cmdResult = cmd.Run()

	cmdStdout = stdoutBuffer.String()
	cmdStderr = stderrBuffer.String()
	cmdRetcode = cmd.ProcessState.ExitCode()

	return cmdStdout, cmdStderr, cmdRetcode, cmdResult

}
