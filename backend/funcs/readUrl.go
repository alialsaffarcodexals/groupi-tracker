package groupie

import (
	"fmt"
	"io"
	"net/http"
)

func ReadUrl(apiURL string) []byte {
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body from %s: %v\n", apiURL, err)
		return nil
	}
	return body
}
