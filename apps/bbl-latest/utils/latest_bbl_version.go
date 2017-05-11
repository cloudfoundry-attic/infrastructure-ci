package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var endpointURL = "https://api.github.com"

func LatestBBLVersion(cachedVersion, cachedLastModified string) (string, string, error) {
	var latestJson struct {
		TagName string `json:"tag_name"`
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repos/cloudfoundry/bosh-bootloader/releases/latest", endpointURL), nil)
	if err != nil {
		// not tested
		return "", "", err
	}
	req.Header.Set("If-Modified-Since", cachedLastModified)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// not tested
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return cachedVersion, cachedLastModified, nil
	}

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// not tested
		return "", "", err
	}

	err = json.Unmarshal(bodyContents, &latestJson)
	if err != nil {
		// not tested
		return "", "", err
	}

	lastModified := resp.Header.Get("Last-Modified")

	return latestJson.TagName, lastModified, nil
}
