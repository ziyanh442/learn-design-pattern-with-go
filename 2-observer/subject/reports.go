package subject

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomTemperature() int {
	return rand.Intn(35-25) + 25
}

func getRandomCrime() (string, string) {
	crimes := []string{
		"Larceny", "Arson", "Murder", "Robbery", "Theft", "Fraud", "Domestic abuse",
	}
	streets := []string{
		"Hampden Meadows", "Humber Spinney", "Hovendens", "Henderson Terrace", "Hurst Cloisters", "Houghton Crescent", "Hill End Farm Lane", "Hermitage Ground", "Harbottle Crescent", "Holly Birches",
	}

	return crimes[rand.Intn(len(crimes))], streets[rand.Intn(len(streets))]
}

func (s *Subject) ReportWeather(duration time.Duration) {
	start := time.Now()

	for {

		s.NotifySubscribers("weather", fmt.Sprintf("Today's temperature is %d°C.", getRandomTemperature()))

		if time.Since(start) >= duration {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *Subject) ReportCrime(duration time.Duration) {
	start := time.Now()

	for {

		crime, location := getRandomCrime()
		s.NotifySubscribers("crime", fmt.Sprintf("%s was committed at %s, more on our website.", crime, location))

		if time.Since(start) >= duration {
			break
		}
		time.Sleep(5 * time.Second)
	}
}
