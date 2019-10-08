package functions

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/MySmartFarm/mysmartfarm_api/constants"
)

func NotifyToLine(msg string) {
	accessToken := os.Getenv(constants.NotifyLineToken)
	if accessToken != "" {
		URL := "https://notify-api.line.me/api/notify"

		u, err := url.ParseRequestURI(URL)
		if err != nil {
			log.Fatal(err)
		}

		c := &http.Client{}

		form := url.Values{}
		form.Add("message", msg)

		body := strings.NewReader(form.Encode())

		req, err := http.NewRequest("POST", u.String(), body)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "Bearer "+accessToken)

		_, err1 := c.Do(req)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
}
