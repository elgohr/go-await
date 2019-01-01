package await_test

import (
	"github.com/elgohr/golang-await/await"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestReturnsResolvedObject(t *testing.T) {
	awaiting := make(chan interface{}, 1)
	expected := "test"
	awaiting <- expected
	returns := await.Await(awaiting, 1*time.Nanosecond)
	if returns != expected {
		t.Errorf("Didn't return %v, but %v", expected, returns)
	}
}

func TestReturnsErrorWhenRunningIntoTimeout(t *testing.T) {
	awaiting := make(chan interface{}, 1)
	returns := await.Await(awaiting, 1*time.Nanosecond)
	if returns.(error).Error() != "Timed out after 1ns" {
		t.Errorf("Didn't return the expected error, but %v", returns)
	}
}

func TestExample(t *testing.T) {
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
	returns := await.Await(awaiting, 1*time.Second)
	if returns != "GET" {
		t.Errorf("Expected GET, but got %v", returns)
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
	await.Await(awaiting, 1*time.Second)
	// Output: GET
}
