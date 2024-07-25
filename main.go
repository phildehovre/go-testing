package main

import (
	"os"
	"time"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency.MyGreeterHandler)))
	sleeper := &ConfigurableSleeper{duration: 500 * time.Millisecond, sleep: time.Sleep}
	// sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
