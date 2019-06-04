package wait

import (
	"errors"
	"fmt"
	"time"
)

func For(goal chan interface{}, timeout time.Duration) interface{} {
	go func() {
		time.Sleep(timeout)
		goal <- errors.New(fmt.Sprintf("Timed out after %v", timeout))
	}()
	return <-goal
}
