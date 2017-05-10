package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Printf("Starting server on port: %s...\n", port)

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.URL.Path {
		case "/latest":
			queryParams := req.URL.Query()
			os := queryParams.Get("os")

			latestBBLVersion, err := getLatestBBLVersion()
			if err != nil {
				// not tested
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Failed to query github.com for latest bbl version. Try again later."))
				fmt.Println(err.Error())
			}

			osTypes := map[string]string{
				"linux": "linux_x86-64",
				"osx":   "osx",
			}
			osType, ok := osTypes[os]
			if ok {
				redirectURL := fmt.Sprintf("https://github.com/cloudfoundry/bosh-bootloader/releases/download/%[1]s/bbl-%[1]s_%s", latestBBLVersion, osType)
				http.Redirect(w, req, redirectURL, http.StatusFound)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("missing required query parameter: os=[osx,linux]"))
			}
		default:
			w.WriteHeader(http.StatusOK)
		}
	}))
	if err != nil {
		log.Fatal(err)
	}
}

func getLatestBBLVersion() (string, error) {
	var latestJson struct {
		TagName string `json:"tag_name"`
	}

	response, err := http.Get("https://api.github.com/repos/cloudfoundry/bosh-bootloader/releases/latest")
	if err != nil {
		// not tested
		return "", err
	}
	defer response.Body.Close()

	bodyContents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// not tested
		return "", err
	}

	err = json.Unmarshal(bodyContents, &latestJson)
	if err != nil {
		// not tested
		return "", err
	}

	return latestJson.TagName, nil
}
