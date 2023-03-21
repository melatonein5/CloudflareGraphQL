package CloudflareHTTP

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//Base API URL (needs to be exportable)
const UrlBase string = "https://api.cloudflare.com/client/v4/graphql"

//Private functions

//getReq() is a private function which will perform a simple get request
func getReq(urlPath, authToken string) (string, error) {
	method := "GET"

	url := UrlBase + urlPath

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", authToken)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//postReq is a private function which performs a simple post request
func postReq(urlPath, authToken, payload string) (string, error) {
	method := "POST"

	url := UrlBase + urlPath
	client := &http.Client{}
	//Convert the string payload into a ioreader
	reqBody := strings.NewReader(payload)

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", authToken)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
