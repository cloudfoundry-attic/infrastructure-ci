package utils

func SetEndpoint(url string) {
	endpointURL = url
}

func ResetEndpoint() {
	endpointURL = "https://api.github.com"
}
