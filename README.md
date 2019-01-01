# golang-await
Test asynchronous behaviour in Golang

`go get -u github.com/elgohr/golang-await`

## Chuck Norris doesn't sleep, he waits
I saw so many test code, where people are waiting on async execution by sleeping (time.Sleep).  
This is
* not efficient, as you may sleep longer for checking values then you would need
* not consistent , as you may sleep longer than your timeout
* dangerous, as some code didn't even have timeouts

In this way I'm trying to illustrate a way to do this easily, by using standard goroutines and channels.  
You may also use this as a library for not playing copy cat :-)

## Example
```
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
```
