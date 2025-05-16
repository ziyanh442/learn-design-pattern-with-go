package main

import (
	"2observer/observer"
	"2observer/subject"
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Starting...")

	subject := subject.NewSubject()

	observerOne := observer.NewObserver("Huckleberry News Outlet")
	observerTwo := observer.NewObserver("New Bork Post")

	subject.Register(observerOne, "weather")
	subject.Register(observerTwo, "weather")
	subject.Register(observerTwo, "crime")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		subject.ReportWeather(1 * time.Minute)
		wg.Done()
	}()

	go func() {
		subject.ReportCrime(1 * time.Minute)
		wg.Done()
	}()

	time.Sleep(25 * time.Second)

	subject.Deregister(observerTwo, "crime")

	time.Sleep(10 * time.Second)

	subject.Register(observerOne, "crime")

	wg.Wait()

	fmt.Println("Stopping...")

}
