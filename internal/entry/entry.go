// Copyright 2026 Vacui Development Team - Apache License 2.0
// https://github.com/vacuiteam/hwapp

package entry

import "sync"

// Entry represents the Vacui hardware application entry.
type Entry struct {
	mu sync.Mutex
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
	return nil
}

func (e *Entry) close() {

}
