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

package linebreakconverter

import (
	"io"

	"golang.org/x/text/transform"
)

// ConvertLinebreakToWindows converts
//
//   - Windows: \r\n (CRLF)
//   - Mac: \r (CR)
//   - Unix/Linux: \n (LF)
//
// into
//
//   - Windows: \r\n (CRLF)
//
// linebreaks
func ConvertLinebreakToWindows(input string) (result string, n int, err error) {

	// *** Bugfix ***
	//
	//	Problem:
	//
	//		convertLinebreakToWindows can not handle input strings that contains
	//			- Windows: \r\n (CRLF)
	//			- Mac: \r (CR)
	// 		linebreaks.
	//
	// Solution (not nice but working):
	//
	//		1. convert to LF (Unix/Linux) to remove CRLF and CR
	//		2. convert to CRLF (Windows)

	//return transform.String(new(convertLinebreakToWindows), input)

	temp, n, err := transform.String(new(convertLinebreakToUnix), input)
	if err != nil {
		return temp, n, err
	}
	return transform.String(new(convertLinebreakToWindows), temp)

}

// ConvertLinebreakToWindowsReader returns an io.Reader that converts
//
//   - Windows: \r\n (CRLF)
//   - Mac: \r (CR)
//   - Unix/Linux: \n (LF)
//
// into
//
//   - Windows: \r\n (CRLF)
//
// linebreaks
func ConvertLinebreakToWindowsReader(r io.Reader) io.Reader {

	// *** Bugfix ***
	//
	//	Problem:
	//
	//		convertLinebreakToWindows can not handle input strings that contains
	//			- Windows: \r\n (CRLF)
	//			- Mac: \r (CR)
	// 		linebreaks.
	//
	// Solution (not nice but working):
	//
	//		1. convert to LF (Unix/Linux) to remove CRLF and CR
	//		2. convert to CRLF (Windows)

	//return transform.NewReader(
	//	r,
	//	new(convertLinebreakToWindows),
	//)

	return transform.NewReader(
		transform.NewReader(
			r,
			new(convertLinebreakToWindows),
		),
		new(convertLinebreakToUnix),
	)

}

// ConvertLinebreakToWindowsWriter returns an io.Writer that converts
//
//   - Windows: \r\n (CRLF)
//   - Mac: \r (CR)
//   - Unix/Linux: \n (LF)
//
// into
//
//   - Windows: \r\n (CRLF)
//
// linebreaks
func ConvertLinebreakToWindowsWriter(r io.Writer) io.Writer {

	// *** Bugfix ***
	//
	//	Problem:
	//
	//		convertLinebreakToWindows can not handle input strings that contains
	//			- Windows: \r\n (CRLF)
	//			- Mac: \r (CR)
	// 		linebreaks.
	//
	// Solution (not nice but working):
	//
	//		1. convert to LF (Unix/Linux) to remove CRLF and CR
	//		2. convert to CRLF (Windows)

	//return transform.NewWriter(
	//	r,
	//	new(convertLinebreakToWindows),
	//)

	return transform.NewWriter(
		transform.NewWriter(
			r,
			new(convertLinebreakToWindows),
		),
		new(convertLinebreakToUnix),
	)

}

type convertLinebreakToWindows struct {
	previous byte
	current  byte
	next     byte
	cr       byte
	lf       byte
}

func (n *convertLinebreakToWindows) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {

	//  Transformer interface (golang.org/x/text/transform):
	//
	//		type Transformer interface {
	//			Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
	//			Reset()
	//		}

	for nDst < len(dst) && nSrc < len(src) {
		n.current = src[nSrc]
		if n.current == n.lf {
			if nDst+1 == len(dst) {
				break
			}
			dst[nDst] = n.cr
			dst[nDst+1] = n.lf
			nSrc++
			nDst += 2
		} else {
			dst[nDst] = n.current
			nSrc++
			nDst++
		}
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

func (n *convertLinebreakToWindows) Reset() {

	//  Transformer interface (golang.org/x/text/transform):
	//
	//		type Transformer interface {
	//			Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
	//			Reset()
	//		}

	n.previous = 0
	n.current = 0
	n.next = 0
	n.cr = byte(13) // --> \r
	n.lf = byte(10) // --> \n

}
