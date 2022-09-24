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

import "testing"

func TestConvertLinebreakToWindows(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		// TODO: Add test cases.
		{"Test---001", "", ""},

		{"Test---002", "\r\n", "\r\n"},
		{"Test---003", "\r", "\r\n"},
		{"Test---004", "\n", "\r\n"},

		{"Test---005", "hello world\r\n", "hello world\r\n"},
		{"Test---006", "hello world\r", "hello world\r\n"},
		{"Test---007", "hello world\n", "hello world\r\n"},

		{"Test---008", "\r\nhello world", "\r\nhello world"},
		{"Test---009", "\rhello world", "\r\nhello world"},
		{"Test---010", "\nhello world", "\r\nhello world"},

		{"Test---011", "hello\r\nworld", "hello\r\nworld"},
		{"Test---012", "hello\rworld", "hello\r\nworld"},
		{"Test---013", "hello\nworld", "hello\r\nworld"},

		{"Test---015", "hello,\r\n\r\nworld", "hello,\r\n\r\nworld"},
		{"Test---016", "hello,\r\rworld", "hello,\r\n\r\nworld"},
		{"Test---017", "hello,\n\nworld", "hello,\r\n\r\nworld"},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got, _, err := ConvertLinebreakToWindows(tt.in)

			if err != nil {
				t.Errorf("\ntransforming error\n   in: %q\n   err: %v", tt.in, err)
				return
			}

			if got != tt.want {
				t.Errorf("\ntransforming got <> want\n   in: %q\n   got: %q\n   want: %q", tt.in, got, tt.want)
			}

		})
	}
}
