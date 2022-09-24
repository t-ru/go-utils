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
```

<b>Example:</b>

```go
cmd := exec.Command("wsl.exe", "--exec",  "sh", "-c", "whoami; whoami; whoamiERRRR; whoami; whoami")
cmd.Stdin = os.Stdin
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
result := cmd.Run()
```

Result: Stdout and Stderr will be displayed correctly. Linebreaks will converted from LF to CRLF<br>

```go
cmd := exec.Command("wsl.exe", "--exec",  "sh", "-c", "whoami; whoami; whoamiERRRR; whoami; whoami")
cmd.Stdin = os.Stdin
cmd.Stdout = io.MultiWriter(os.Stdout, &stderrBuffer)
cmd.Stderr = io.MultiWriter(os.Stderr, &stdoutBuffer)
result := cmd.Run()
```
Result: Stdout and Stderr will not be displayed correctly. Stdout, Stderr, stdoutBuffer, stderrBuffer containing LF.

```go
cmd := exec.Command("wsl.exe", "--exec",  "sh", "-c", "whoami; whoami; whoamiERRRR; whoami; whoami")
cmd.Stdin = os.Stdin
cmd.Stdout = io.MultiWriter(linebreakconverter.ConvertLinebreakToWindowsWriter(os.Stdout), &stdoutBuffer)
cmd.Stderr = io.MultiWriter(linebreakconverter.ConvertLinebreakToWindowsWriter(os.Stderr), &stderrBuffer)
result := cmd.Run()
```
Result: Stdout and Stderr will displayed correctly. Stdout and Stderr containig CRLF. stdoutBuffer and stderrBuffer containing LF

```go
cmd := exec.Command("wsl.exe", "--exec",  "sh", "-c", "whoami; whoami; whoamiERRRR; whoami; whoami")
cmd.Stdin = os.Stdin
cmd.Stdout = linebreakconverter.ConvertLinebreakToWindowsWriter(io.MultiWriter(os.Stdout, &stdoutBuffer))
cmd.Stderr = linebreakconverter.ConvertLinebreakToWindowsWriter(io.MultiWriter(os.Stderr, &stderrBuffer))
result := cmd.Run()
```
Result: Stdout and Stderr will displayed correctly. Stdout, Stderr, stdoutBuffer and stderrBuffer containing CRLF.










