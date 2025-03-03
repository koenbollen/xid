package main

import (
	"os"
	"strconv"

	"github.com/rs/xid"
)

func main() {
	count := 1
	if len(os.Args) > 1 {
		if c, err := strconv.Atoi(os.Args[1]); err == nil {
			count = c
		}
	}
	for range count {
		guid := xid.New()
		os.Stdout.WriteString(guid.String() + "\n")
	}
	os.Stdout.Sync()
}
