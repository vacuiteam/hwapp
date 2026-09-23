// Copyright 2026 Vacui Development Team - Apache License 2.0
// https://github.com/vacuiteam/hwapp

package main

import (
	"fmt"
	"os"

	"github.com/vacuiteam/hwapp/internal/entry"
)

func main() {
	e := entry.New()
	if err := e.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
