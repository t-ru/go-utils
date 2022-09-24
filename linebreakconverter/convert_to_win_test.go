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
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantN      int
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotN, err := ConvertLinebreakToWindows(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertLinebreakToWindows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("ConvertLinebreakToWindows() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotN != tt.wantN {
				t.Errorf("ConvertLinebreakToWindows() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
