package gonexus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API struct
type API struct {
	baseurl  string
	username string
	password string
	client   *http.Client
}

// NewAPI func
func NewAPI(url string, username string, password string) *API {
	return &API{url, username, password, &http.Client{}}
}

// HTTPDo func
func (api *API) HTTPDo(url string, method string) (body []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}
	req.SetBasicAuth(api.username, api.password)
	response, err := api.client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return

}

// GetSearch func
func (api *API) GetSearch(body []byte) (data Search, err error) {
	json.Unmarshal(body, &data)
	return data, nil
}

// GetComponent func
func (api *API) GetComponent(body []byte) (data Component, err error) {
	json.Unmarshal(body, &data)
	return data, nil
}

// APISearch func
func (api *API) APISearch(repository string, search string, ctoken string) (data Search, err error) {
	var url string
	var query string
	if search == "" {
		query = query + ""
		url = fmt.Sprintf("%s/beta/search?repository=%s", api.baseurl, repository)
	} else {
		query = query + "&q=" + search
	}
	if ctoken == "" {
		query = query + ""
	} else {
		query = query + "&continuationToken=" + ctoken
	}
	url = fmt.Sprintf("%s/beta/search?repository=%s%s", api.baseurl, repository, query)
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	data, err = api.GetSearch(body)
	return data, err
}

// APIComponent func
func (api *API) APIComponent(id string) (data Component, err error) {
	var url string
	url = api.baseurl + "/beta/components/" + id
	body, err := api.HTTPDo(url, "GET")
	if err != nil {
		return
	}
	data, err = api.GetComponent(body)
	return data, err
}

// APIDelComponent func
func (api *API) APIDelComponent(id string) error {
	var url string
	url = api.baseurl + "/beta/components/" + id
	_, err := api.HTTPDo(url, "DELETE")
	return err
}
