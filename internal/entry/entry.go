// Copyright 2026 Vacui Development Team - Apache License 2.0
// https://github.com/vacuiteam/hwapp

package entry

import (
	"fmt"
	"sync"

	"github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

// Entry represents the Vacui hardware application entry.
type Entry struct {
	mu sync.Mutex
	uc unicorn.Unicorn
}

func New() *Entry {
	return &Entry{}
}

// Start initializes hardware application's resources and closes them on the exit.
func (e *Entry) Start() error {
	e.mu.Lock()
	defer e.mu.Unlock()
	err := e.init()
	if err != nil {
		return err
	}
	defer e.close()
	return nil
}

func (e *Entry) init() error {
	uc, err := unicorn.NewUnicorn(unicorn.ARCH_X86, unicorn.MODE_32)
	if err != nil {
		return fmt.Errorf("initializing unicorn.Unicorn failed: %v", err)
	}
	e.uc = uc
	return nil
}

func (e *Entry) close() {
	if e.uc != nil {
		e.uc.Close()
	}
}
