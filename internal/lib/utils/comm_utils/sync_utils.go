package comm_utils

import (
	"sync"
	"time"
)

// WaitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

// WaitError waits for the waitgroup. Group is stopped immediately if any error occurs.
// Returns true if waiting timed out.
func WaitError(wg *sync.WaitGroup, errC chan error) error {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return nil // completed normally
	case e := <-errC:
		return e // error
	}
}
