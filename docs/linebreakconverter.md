# Package linebreakconverter

## Source

- [linebreakconverter](../pkg/linebreakconverter/)

## Usage
```go
import (
    "github.com/t-ru/go-utils/pkg/linebreakconverter"
)
```

## Documentation

### <span id="ConvertLinebreakToUnix">ConvertLinebreakToUnix</span>
Converts linebreaks (Windows CRLF \r\n , Mac CR \r, Unix/Linux LF \n) to Unix/Linux LF \n linebreaks<br>

<b>Signature:</b>

```go
func ConvertLinebreakToUnix(input string) (result string, n int, err error)
```
<b>Example:</b>

```go
package main

import (
	"fmt"

	"github.com/t-ru/go-utils/pkg/linebreakconverter"
)

func main() {
	in := "\r\nhello\r\n\r\nworld\r\n"
	result, result_len, result_err := linebreakconverter.ConvertLinebreakToUnix(in)
	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", result_len)
	fmt.Printf("%#v\n", result_err)
}
```

### <span id="ConvertLinebreakToWindows">ConvertLinebreakToWindows</span>
Converts linebreaks (Windows CRLF \r\n , Mac CR \r, Unix/Linux LF \n) to Windows LF \r\n linebreaks<br>

<b>Signature:</b>

```go
func ConvertLinebreakToWindows(input string) (result string, n int, err error)
```
<b>Example:</b>

```go
package main

import (
	"fmt"

	"github.com/t-ru/go-utils/pkg/linebreakconverter"
)

func main() {
	in := "\n\nhellon\n\nworld\n\n"
	result, result_len, result_err := linebreakconverter.ConvertLinebreakToWindows(in)
	fmt.Printf("%#v\n", result)
	fmt.Printf("%#v\n", result_len)
	fmt.Printf("%#v\n", result_err)
}
```

### <span id="ConvertLinebreakToUnixReader">ConvertLinebreakToUnixReader</span>
Returns an io.Reader that converts linebreaks (Windows CRLF \r\n , Mac CR \r, Unix/Linux LF \n) to Unix/Linux LF \n linebreaks<br>

<b>Signature:</b>

```go
func ConvertLinebreakToUnixReader(r io.Reader) io.Reader
```

### <span id="ConvertLinebreakToUnixWriter">ConvertLinebreakToUnixWriter</span>
Returns an io.Writer that converts linebreaks (Windows CRLF \r\n , Mac CR \r, Unix/Linux LF \n) to Unix/Linux LF \n linebreaks<br>

<b>Signature:</b>

```go
func ConvertLinebreakToUnixWriter(r io.Writer) io.Writer
```

### <span id="ConvertLinebreakToWindowsReader">ConvertLinebreakToWindowsReader</span>
Returns an io.Reader that converts linebreaks (Windows CRLF \r\n , Mac CR \r, Unix/Linux LF \n) to Windows LF \r\n linebreaks<br>

<b>Signature:</b>

```go
func ConvertLinebreakToWindowsReader(r io.Reader) io.Reader
```

### <span id="ConvertLinebreakToWindowsWriter">ConvertLinebreakToWindowsWriter</span>
Returns an io.Writer that converts linebreaks (Windows CRLF \r\n , Mac CR \r, Unix/Linux LF \n) to Windows LF \r\n linebreaks<br>

<b>Signature:</b>

```go
func ConvertLinebreakToWindowsWriter(r io.Writer) io.Writer