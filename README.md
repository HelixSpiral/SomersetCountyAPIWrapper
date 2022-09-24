Somerset County API Wrapper
---

This is a simple wrapper for the Somerset County API provided here: https://api.somersetcounty-me.org/

Currently only the `DispatchLog` endpoint is working so that's all this wrapper provides.

This API also has some oddities, like returning values having random whitespace characters in them. This package trims all results from the API before returning the results.

If you need the values with the whitespaces not removed, fork the project and look at `structInternal.go`. You should be able to remove the custom unmarshaler and get the whitespaces.

Usage:
---

```go
package main

import (
	"log"
	"time"

	somersetcountywrapper "github.com/HelixSpiral/SomersetCountyAPIWrapper"
)

func main() {
	currentDate := time.Now()

	sw := somersetcountywrapper.NewWrapper()

	logs, err := sw.GetDispatch(currentDate.Format("20060102"))
	if err != nil {
		panic(err)
	}

	for x, y := range logs {
		log.Printf("[%d]: %+v\r\n", x, y)
	}
}

```