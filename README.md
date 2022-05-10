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
go get github.com/colored-paper/epsilon
```
2. Add import `github.com/colored-paper/epsilon` to your source code.
```go
package main

import (
	"time"
	
	"github.com/colored-paper/epsilon"
)

func main() {
	e := epsilon.New(time.Now(), pid)
	id, err := e.Next()
}
```