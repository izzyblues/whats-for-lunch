package main

import (
	"log"

	"time"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
	"github.com/izzyblues/whats-for-lunch/lunch"
	"github.com/izzyblues/whats-for-lunch/weather"
	"golang.org/x/net/context"
)

const owmAppID = ""

func main() {
	ctx := context.Background()
	projectID := ""

	ws := weather.NewWeatherService(owmAppID)
	wi, err := ws.CurrentWeatherInfo()
	if err != nil {
		log.Fatalf("Failed to get current weather %v", err)
	}

	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	kind := "lunch"

	lunchKey := datastore.NameKey(kind, uuid.New().String(), nil)

	l := lunch.Lunch{
		Date:        time.Now(),
		WeatherInfo: wi,
	}

	_, err = client.Put(ctx, lunchKey, &l)
	if err != nil {
		log.Fatalf("Failed to put entity %v", err)
	}
}
