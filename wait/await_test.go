package wait_test

import (
	"fmt"
	"github.com/elgohr/go-await/wait"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReturnsResolvedObject(t *testing.T) {
	awaiting := make(chan interface{}, 1)
	expected := "test"
	awaiting <- expected
	returns := wait.For(awaiting, 1*time.Nanosecond)
	if returns != expected {
		t.Errorf("Didn't return %v, but %v", expected, returns)
	}
}

func TestReturnsErrorWhenRunningIntoTimeoutWithoutAnswer(t *testing.T) {
	awaiting := make(chan interface{}, 1)
	returns := wait.For(awaiting, 1*time.Nanosecond)
	if returns.(error).Error() != "Timed out after 1ns" {
		t.Errorf("Didn't return the expected error, but %v", returns)
	}
}

func TestReturnsErrorWhenRunningIntoTimeoutWithLaterAnswer(t *testing.T) {
	awaiting := make(chan interface{}, 1)
	returns := wait.For(awaiting, 1*time.Nanosecond)
	go func() {
		time.Sleep(2 * time.Nanosecond)
		awaiting <- "test"
	}()
	if returns.(error).Error() != "Timed out after 1ns" {
		t.Errorf("Didn't return the expected error, but %v", returns)
	}
}

func ExampleAwait() {
	awaiting := make(chan interface{}, 1)
	remoteServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		awaiting <- r.Method
		w.WriteHeader(http.StatusOK)
	}))
	defer remoteServer.Close()

	thisCouldBeYourAsyncFunction := func() {
		go func() {
			res, _ := http.Get(remoteServer.URL)
			defer res.Body.Close()
		}()
	}
	thisCouldBeYourAsyncFunction()
	fmt.Println(wait.For(awaiting, 1*time.Second))
	// Output: GET
}
