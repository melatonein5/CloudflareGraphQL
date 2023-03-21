package CloudflareHTTP

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//Base API URL (needs to be exportable)
const UrlBase string = "https://api.cloudflare.com/client/v4/graphql"

//Structure to make the functions Exportable
//Exportable public functions

type API struct {
	UrlPath   string
	AuthToken string
	Payload   string
}

//Exportable functions

//Get() is a function which will perform a simple get request
func (api API) Get() (string, error) {
	//Define the HTTP method to use
	method := "GET"

	//Define the full URL
	url := UrlBase + api.UrlPath

	//Create the HTTP client and request, and check and return errors
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	//Add the API header as an authorization header
	req.Header.Add("Authorization", api.AuthToken)

	//Execute the request and check for errors
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	//Return as string anc check for errors
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

//Post is a private function which performs a simple post request
func (api API) Post() (string, error) {
	//Define the HTTP method to use
	method := "POST"

	//Define the full URL
	url := UrlBase + api.UrlPath

	//Create the HTTP client and request, and check for errors
	client := &http.Client{}
	//Convert the string payload into a ioreader
	reqBody := strings.NewReader(api.Payload)
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return "", err
	}
	//Add the API header as an authorization header
	req.Header.Add("Authorization", api.AuthToken)

	//Execute the request and check for errors
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	//Return as string anc check for errors
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
