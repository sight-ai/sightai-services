package comm_utils

import (
	"errors"
	"net/http"
)

func GetRedirectedUrl(keyUrl string) string {
	req, err := http.NewRequest("GET", keyUrl, nil)
	if err != nil {
		panic(err)
	}
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}

	response, err := client.Do(req)
	if response != nil && err != nil {
		if response.StatusCode == http.StatusFound ||
			response.StatusCode == http.StatusMovedPermanently ||
			response.StatusCode == http.StatusTemporaryRedirect ||
			response.StatusCode == http.StatusPermanentRedirect { //status code 302
			url, err := response.Location()
			if err != nil {
				return keyUrl
			}
			return url.String()
		}
	}
	return keyUrl
}
