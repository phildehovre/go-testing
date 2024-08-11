package main

func main() {
	// log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependency.MyGreeterHandler)))
	// sleeper := &ConfigurableSleeper{duration: 500 * time.Millisecond, sleep: time.Sleep}
	// sleeper := &DefaultSleeper{}
	// Countdown(os.Stdout, sleeper)

	// startNow := time.Now()
	// cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	// ch := make(chan string)
	// var wg sync.WaitGroup

	// for _, city := range cities {
	// 	wg.Add(1)
	// 	go FetchWeather(city, ch, &wg)
	// }
	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for result := range ch {
	// 	fmt.Println(result)
	// }

	// fmt.Println("this operation took:", time.Since(startNow))
}
