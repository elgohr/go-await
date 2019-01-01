# golang-await
Test asynchronous behaviour in Golang

`go get -u github.com/elgohr/golang-await`

## Chuck Norris doesn't sleep, he waits
I did see so many test code, where people are waiting on async execution by sleeping (time.Sleep).  
This is
* not efficient, as you may sleep longer for checking values then you would need
* not consistent , as you may sleep longer than your timeout
* dangerous, as some code didn't even have timeouts

In this way I'm trying to illustrate a way to do this easily, by using standard goroutines and channels.
You may also use this as a library for not playing copy cat :-)
