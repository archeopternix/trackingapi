// encapsulates the low level eTracker REST API calls and returns JSON
// messages that will be parsed by methods in tracking.go
//
// eTracker REST API documentation:
// [eTracker developer API documentation](https://www.etracker.com/en/docs/integration-setup-2/developer-apis/management-of-authorisation-profiles-an-sub-users/)
//
// Manage sub-users
// The setup of the general route for sub-users is as follows:
//    `https://ws.etracker.com/api/v6/subuserPublic/user`
//
// Query sub-users
// The following Curl example shows a GET request for querying already existing sub-users:
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/user' -X GET -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo'`
//
// As a response to the GET request, you will then receive:
//     `[[{"id":"6","name":"Dalton","fname":"Lars","version":"1","subid":"2","role":"Read permissions","role_type":"read_only","enable":"1","login":"3","role_id":"5"},`
//     `{"id":"8","name":"Write permission","fname":"Lars","version":"1","subid":"3","role":`
//     `"Read permissions","role_type":"read_only","enable":"1","login":"0","role_id":"5"},`
//     `{"id":"9","name":"Read-Write permission","fname":"Lars","version":"1","subid":"4",`
//     `"role":"Read and write permissions","role_type":"read_only","enable":"1","login":"3",`
//     `"role_id":"7"}]]`
//
// Query a single sub-user
// The following Curl example shows a GET request for querying one single existing sub-users:
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/user/6' -X GET -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo'`
//
// As a response to the GET request, you will then receive:
//     `[{"id":"6","name":"Dalton","fname":"Lars","version":"1","subid":"2","role":"Read permissions","role_type":"read_only","enable":"1","login":"3","role_id":"5"}]`
//
// Create sub-user
// The following Curl example shows you the setup of a sub-user:
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/user' -X POST -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo'--data 'enable=1&amp;sex=0&amp;fname=John&amp;name=Doe&amp;email=qa%40etracker.com&amp;role_id=7&amp;language=de&amp;pass=test'`
//
// The information contained in the “–data” parameter is:
// * enable = active, value 1 or inactive, value 0
// * sex = female, value 0; sex male, value 1
// * fname = First name
// * name = Last name
// * email = Email address
// * role_id = Identifier of the user profile
// * pass = Password
// * language = language (Spanish = es, English = en, French = fr and German = de) In the “–data” parameter, the following information must always be contained: role_id, name, fname, sex, pass, email, language, enable
//
// _Note: The available “roleIDs” can be queried with a GET request (see “Query profile“)._
//
// Delete sub-user
// The route for deleting a sub-user is as follows:
//     `https://ws.etracker.com/api/v6/subuserPublic/user/<userId>`
//
// The following Curl example shows a DELETE request which deletes the sub-user with the ID 8:
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/user/8' -X DELETE -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo'`
//
// _Note: The available “userIDs” can be queried with a GET request (see “Query sub-users“)._
//
// Managing authorisation profiles
// The setup of the general route for the authorisation profiles is as follows:
//     `https://ws.etracker.com/api/v6/subuserPublic/role`
//
// Query profile
// You can set a GET request in order to query existing authorisation profiles in your account. The following Curl example is good for demonstration purposes here:
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/role' -X GET -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo'`
//
// As a response to the GET request, you will then receive:
//     `[[{"id":"7","name":"Read and write `
//     `permissions","version":"1","users":"1"},{"id":"5","name":"Read `
//     `permissions","version":"1","users":"2"}]]`
//
// Create profile
// The following Curl example shows you how to create an authorisation profile that grants access to two sub-users:
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/role' -X POST -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo' --data 'name=Test&amp;role_type=admin&amp;multi_client_access_mode=custom&amp;multi_clients_selection%5B%5D=18855&amp;multi_clients_selection%5B%5D=260960'`
//
// The information contained in the “–data” parameter is:
// * name = The name for the authorisation profile can be freely chosen and must be unique
// * role_type = Read and write permissions or just read permissions, value can be: read_only or admin
// * multi_client_access_mode = Set whether or not multiple clients can be accessed using this authorisation profile. Values can be: none, all or custom.
// * multi_clients_selection = Details of the multiple clients to whom access is granted. More than just one can be entered.
//
// _Important note: In the “–data” parameter, the following information must always be contained: name, role_type and multi_client_access_mode If the value “custom” is assigned to multi_client_access_mode, the corresponding accounts (multiple clients) must be set via multi_clients_selection (see Curl example)._
//
// If the value “custom” is assigned to multi_client_access_mode, the corresponding accounts (multiple clients) must be set via multi_clients_selection (see Curl example).
//
// Delete profile
// The setup of the route for deleting an authorisation profile is as follows:
//     `https://ws.etracker.com/api/v6/subuserPublic/role/<roleId>`

// The following Curl shows a DELETE request, which deletes the profile with the ID 7.
//     `curl 'https://ws.etracker.com/api/v6/subuserPublic/role/7' -X DELETE -H 'X-ET-email: qa@etracker.com' -H 'X-ET-developerToken: ab7891ca89d9b4d10dc1703a7f0214256babe6c9' -H 'X-ET-accountId: 18854' -H 'X-ET-password: demo'`
//
// _Note: The available “roleIDs” can be queried with a GET request (see Curl example in “Query profile“). It will not be possible to delete an authorisation profile if it is also assigned to another user._

package trackingapi

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
	conf := Config{Email: "etracker@osram.info", Token: "", AccountID: id, Password: password}
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
