// etrackerapi.go encapsulates the low level eTracker REST API calls and returns JSON
// messages that will be parsed by methods in tracking.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Config holds the credentials for authenticating against eTracker API
type Config struct {
	AccountID string // eTracker Account ID
	Password  string // Password for the "admin" account
	Email     string // "admin" eMail address
	Token     string // API token to be requested by eTracker
}

// NewConfig creates new account configuration with preset email and token
func NewConfig(id string, password string) *Config {
	conf := Config{Email: "etracker@osram.info", Token: "6638704be30b2711dd9ba5b7b5b7f3e7b16ba8e1", AccountID: id, Password: password}
	return &conf
}

// checkConfig checks if password and token fields are filled
func (c Config) checkConfig() (err error) {
	if len(c.Password) < 1 || len(c.Token) < 1 {
		return fmt.Errorf("Error: Password or Token not set in 'Config'")
	}
	return nil
}

// ETrackerAPI encapsulates the eTracker REST calls and provides back a JSON answer
// implements the IJsonApi interface
type ETrackerAPI struct {
	config *Config
}

//NewETrackerAPI returns an new JSON eTracker API
func NewETrackerAPI(c *Config) *ETrackerAPI {
	m := new(ETrackerAPI)
	m.config = c
	return m
}

// eTracker GET REST call with JSON response for one user by id for an account
func (m ETrackerAPI) getUserByIDJSON(id string) (json []byte, err error) {
	userjson, err := m.eTrackerGET("user/" + id)
	if err != nil {
		return nil, err
	}
	return userjson, nil
}

// eTracker GET REST call with JSON response for all users for an account
func (m ETrackerAPI) getUsersJSON() (json []byte, err error) {
	userjson, err := m.eTrackerGET("user")
	if err != nil {
		return nil, err
	}
	return userjson, nil
}

// eTracker GET REST call with JSON response for all roles for an account
func (m ETrackerAPI) getRolesJSON() (json []byte, err error) {
	rolejson, err := m.eTrackerGET("role")
	if err != nil {
		return nil, err
	}
	return rolejson, nil
}

// eTracker POST REST call with JSON response for creation of a new user for an account
func (m ETrackerAPI) createUserJSON(data *url.Values) (json []byte, err error) {
	var req *http.Request
	var body []byte

	req, err = http.NewRequest("POST", "https://ws.etracker.com/api/v6/subuserPublic/user", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// Call the generic eTracker POST, DELETE call
	body, err = m.eTrackerPostDelete(req)
	if err != nil {
		return body, err
	}
	return body, nil
}

// eTracker DELETE REST call with JSON response for creation of a new user for an account
func (m ETrackerAPI) deleteUserJSON(id string) (json []byte, err error) {
	var req *http.Request
	var body []byte

	data := url.Values{}
	req, err = http.NewRequest("DELETE", "https://ws.etracker.com/api/v6/subuserPublic/user/"+id, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	// Call the generic eTracker POST, DELETE call
	body, err = m.eTrackerPostDelete(req)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Generic GET call helper that does all the authentication for us and returns
// the plain JSON response back to the caller
func (m ETrackerAPI) eTrackerGET(call string) (rawjson []byte, err error) {
	var res *http.Response
	var req *http.Request

	err = m.config.checkConfig()
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("GET", "https://ws.etracker.com/api/v6/subuserPublic/"+call, nil)
	if err != nil {
		return nil, err
	}

	// Set header manually otherwise the key is camel-cased
	header := make(http.Header)
	header["X-ET-email"] = []string{m.config.Email}
	header["X-ET-developerToken"] = []string{m.config.Token}
	header["X-ET-accountId"] = []string{m.config.AccountID}
	header["X-ET-password"] = []string{m.config.Password}
	req.Header = header

	client := new(http.Client)
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	rawjson, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(rawjson), "errorCode") {
		return nil, fmt.Errorf("%v", string(rawjson))
	}
	return rawjson, nil
}

// Generic POST and DELETE helper that does all the authentication for us and returns
// the plain JSON body response back to the caller
func (m ETrackerAPI) eTrackerPostDelete(req *http.Request) (body []byte, err error) {
	var res *http.Response

	err = m.config.checkConfig()
	if err != nil {
		return nil, err
	}

	// Set header manually otherwise the key is camel-cased
	header := make(http.Header)
	header["X-ET-email"] = []string{m.config.Email}
	header["X-ET-developerToken"] = []string{m.config.Token}
	header["X-ET-accountId"] = []string{m.config.AccountID}
	header["X-ET-password"] = []string{m.config.Password}

	req.Header = header

	client := new(http.Client)
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("HTTP %v, %w", res.Status, err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP %v, %v", res.Status, string(body))
	}

	return body, nil
}
