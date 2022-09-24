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

package wslterminate

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/t-ru/go-utils/pkg/linebreakconverter"
)

type options struct {
	Distribution string
	StdoutSilent bool
	StderrSilent bool
}

type Option func(f *options)

func Distribution(value string) Option {
	return func(o *options) {
		o.Distribution = value
	}
}

func StdoutSilent(value bool) Option {
	return func(o *options) {
		o.StdoutSilent = value
	}
}

func StderrSilent(value bool) Option {
	return func(o *options) {
		o.StderrSilent = value
	}
}

func Run(opt ...Option) (stdout string, stderr string, retcode int, err error) {

	opts := &options{
		Distribution: "",
		StdoutSilent: false,
		StderrSilent: false,
	}

	for _, applyOpt := range opt {
		applyOpt(opts)
	}

	var stdoutBuffer bytes.Buffer
	var stderrBuffer bytes.Buffer
	var cmd *exec.Cmd
	var cmdArgs []string
	var cmdResult error
	var cmdStdout string
	var cmdStderr string
	var cmdRetcode int

	cmdArgs = append(cmdArgs, "--terminate")

	if opts.Distribution != "" {
		cmdArgs = append(cmdArgs, opts.Distribution)
	}

	cmd = exec.Command("wsl.exe", cmdArgs...)

	cmd.Stdin = os.Stdin

	if opts.StdoutSilent {
		cmd.Stdout = &stdoutBuffer
	} else {
		// os.Stdout = crlf, stdoutBuffer = lf
		cmd.Stdout = io.MultiWriter(linebreakconverter.ConvertLinebreakToWindowsWriter(os.Stdout), &stdoutBuffer)
		// os.Stdout = crlf, stdoutBuffer = crlf
		//cmd.Stdout = linebreakconverter.ConvertLinebreakToWindowsWriter(io.MultiWriter(os.Stdout, &stdoutBuffer))
	}

	if opts.StderrSilent {
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
