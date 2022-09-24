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

package wsl

import (
	"errors"

	"github.com/t-ru/go-utils/pkg/wsl/wslexec"
	"github.com/t-ru/go-utils/pkg/wsl/wslexport"
	"github.com/t-ru/go-utils/pkg/wsl/wslimport"
	"github.com/t-ru/go-utils/pkg/wsl/wslterminate"
	"github.com/t-ru/go-utils/pkg/wsl/wslunregister"
)

func Exec(opt ...wslexec.Option) (stdout string, stderr string, retcode int, err error) {
	return wslexec.Run(opt...)
}

func Import(opt ...wslimport.Option) (stdout string, stderr string, retcode int, err error) {
	return "", "", 1, errors.New("not implemented")
	//return wslimport.Run(opt...)
}

func Export(opt ...wslexport.Option) (stdout string, stderr string, retcode int, err error) {
	return "", "", 1, errors.New("not implemented")
	//return wslexport.Run(opt...)
}

func Terminate(opt ...wslterminate.Option) (stdout string, stderr string, retcode int, err error) {
	return wslterminate.Run(opt...)
}

func Unregister(opt ...wslunregister.Option) (stdout string, stderr string, retcode int, err error) {
	return wslunregister.Run(opt...)
}
