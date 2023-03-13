## **Epsilon** is a 64-bit ID generator similar to [Twitter Snowflake](https://github.com/twitter-archive/snowflake).

### ID is composed of:
  - Timestamp 45bits
    - increases in units of about 100µs(microseconds).
    - can be stored for a total of 68 years.
  - ParentsID 9bits
  - SequenceNumber 10bits

### Generated ID result expressed in bits
```[ Timestamp 0 ~ 2^45 - 1 ][ PID 0 ~ 2^9-1][ Num 0 ~ 2^10-1]```

### How to use
1. Add the epsilon package to your project dependencies (go.mod).
```shell
go get github.com/end-2/epsilon
```
2. Add import `github.com/end-2/epsilon` to your source code.

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/end-2/epsilon"
)

func main() {
	// pid can be from 0 to 2^9-1
	pid := uint32(0)
	e, err := epsilon.New(time.Now(), pid)
	if err != nil {
		log.Fatal(err)
	}
	id, err := e.Next()
	if err != nil {
		fmt.Println(err)
    }
	fmt.Println(id)
}
```

#### If you want to see the form of the generated ID, use "go test -v" command.
