// +build prod

package main

import (
	"fmt"
	"time"
)

func init() {
	version = fmt.Sprintf("production-%s", time.Now())
}
