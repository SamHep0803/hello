package components

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/samhep0803/hello/internal/api"
	"github.com/samhep0803/hello/internal/creds"
	"github.com/spf13/viper"
)

func NewWeatherComponent() *tview.TextView {
	weather := tview.NewTextView()

	weather.SetBorder(true).
		SetTitle("Weather")

	city := viper.GetString("weather.city")

	tokens, err := creds.GetCreds()
	if err != nil {
		return nil
	}

	weatherToken := tokens["weather"]

	forecast, err := api.GetForecast(city, weatherToken)
	if err != nil {
		return nil
	}

	fmt.Fprintf(weather, "📌 Location: %s\n", city)
	fmt.Fprintf(weather, "☁️ Forecast: %s\n", forecast)

	return weather
}
