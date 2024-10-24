package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"telegram-chatbot/models"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetWeather(update tgbotapi.Update) (string, error) {
	q := update.Message.Text

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=f4181eff4691479584d110001242310&q=" + q + "&days=1&aqi=no&alerts=no")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New("API Cuaca tidak tersedia")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var weather models.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return "", err
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	infoPlace := fmt.Sprintf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf(
			"%s - %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
		infoPlace = infoPlace + message
	}
	return infoPlace, nil
}
