package lunch

import (
	"time"

	"github.com/izzyblues/whats-for-lunch/weather"
)

type Lunch struct {
	Date        time.Time
	People      []Person
	Place       string
	WeatherInfo weather.Info
}

type Person struct {
	Name string
}
