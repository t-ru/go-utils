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