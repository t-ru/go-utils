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

// ConvertLinebreakToUnix converts
//
//   - Windows: \r\n (CRLF)
//   - Mac: \r (CR)
//   - Unix/Linux: \n (LF)
//
// into
//
//   - Unix/Linux: \n (LF)
//
// linebreaks
func ConvertLinebreakToUnix(input string) (result string, n int, err error) {
	return transform.String(new(convertLinebreakToUnix), input)
}

// ConvertLinebreakToUnixReader returns an io.Reader that converts
//
//   - Windows: \r\n (CRLF)
//   - Mac: \r (CR)
//   - Unix/Linux: \n (LF)
//
// into
//
//   - Unix/Linux: \n (LF)
//
// linebreaks
func ConvertLinebreakToUnixReader(r io.Reader) io.Reader {
	return transform.NewReader(
		r,
		new(convertLinebreakToUnix),
	)
}

// ConvertLinebreakToUnixWriter returns an io.Writer that converts
//
//   - Windows: \r\n (CRLF)
//   - Mac: \r (CR)
//   - Unix/Linux: \n (LF)
//
// into
//
//   - Unix/Linux: \n (LF)
//
// linebreaks
func ConvertLinebreakToUnixWriter(r io.Writer) io.Writer {
	return transform.NewWriter(
		r,
		new(convertLinebreakToUnix),
	)
}

type convertLinebreakToUnix struct {
	previous byte
	current  byte
	next     byte
	cr       byte
	lf       byte
}

func (n *convertLinebreakToUnix) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {

	//  Transformer interface (golang.org/x/text/transform):
	//
	//		type Transformer interface {
	//			Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
	//			Reset()
	//		}

	for nDst < len(dst) && nSrc < len(src) {
		n.current = src[nSrc]
		switch n.current {
		case n.cr:
			dst[nDst] = n.lf
		case n.lf:
			if n.previous == n.cr {
				nSrc++
				n.previous = n.current
				continue
			}
			dst[nDst] = n.lf
		default:
			dst[nDst] = n.current
		}
		n.previous = n.current
		nDst++
		nSrc++
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return

}

func (n *convertLinebreakToUnix) Reset() {

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
