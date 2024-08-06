package currency

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ResponseData struct {
	Rates map[string]float64 `json:"rates"`
}

func Convert(amount float64, from string, to string) (float64, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	res, err := http.Get("https://openexchangerates.org/api/latest.json?prettyprint=false&app_id=" + os.Getenv("OER_API_KEY"))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	rateFrom := data.Rates[from]
	rateTo := data.Rates[to]

	amount = amount * (rateTo / rateFrom)

	return math.Round(amount*100) / 100, nil
}
