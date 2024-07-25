package main

import (
	"os"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency.MyGreeterHandler)))
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
