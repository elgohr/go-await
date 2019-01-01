package await

import (
	"errors"
	"fmt"
	"time"
)

func Await(waitFor chan interface{}, timeout time.Duration) interface{} {
	go func() {
		time.Sleep(timeout)
		waitFor <- errors.New(fmt.Sprintf("Timed out after %v", timeout))
	}()
	return <-waitFor
}
