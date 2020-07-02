// restapi.go encapsulates the low level eTracler REST calls and returns JSON
// messages that will be parsed by methods in etracker.go
package main

import (
	"net/url"
)

// MockTrackerAPI implements the 'TrackingAPI' interface used to initialize a
// 'TrackingAccount' in tests and examples
type MockTrackerAPI struct {
}

// NewMockTrackerAPI creates and returns a '*MockTrackerAPI'
func NewMockTrackerAPI() *MockTrackerAPI {
	m := new(MockTrackerAPI)
	return m
}

// getUserByIdJSON mocks user by ID implementation
func (m MockTrackerAPI) getUserByIDJSON(id string) (json []byte, err error) {
	json = []byte(`[{"id":"22","language":"en","name":"Anja","fname":"Musterfrau","version":"1","subid":"14","role_id":"21","enable":"1","login":"8","email":"a.m@osram.com","sex":"1","newsletterSubscribed":false}]`)
	return json, nil
}

// getUsersJSON mocks user list implementation
func (m MockTrackerAPI) getUsersJSON() (json []byte, err error) {
	json = []byte(`[[{"id":"22","name":"Anja","fname":"Musterfrau","version":"1","subid":"14","role":"Analyze and Edit Access Rights","role_type":"read_only","enable":"1","login":"8","role_id":"21"},{"id":"25","name":"Bluemchen","fname":"Benjamin","version":"1","subid":"17","role":"Analyze and Edit Access Rights","role_type":"read_only","enable":"1","login":"8","role_id":"21"}]]`)
	return json, nil
}

// getRolesJSON mocks role list implementation
func (m MockTrackerAPI) getRolesJSON() (json []byte, err error) {
	json = []byte(`[[{"id":"39","name":"Analytics only Access Rights","version":"1","users":"2"},{"id":"21","name":"Analyze and Edit Access Rights","version":"1","users":"19"}]]`)
	return json, nil
}

// createUserJSON mocks create user implementation
func (m MockTrackerAPI) createUserJSON(data *url.Values) (json []byte, err error) {
	json = []byte(`[{"id":78}]`)
	return json, nil
}

// deleteUserJSON mocks delete user implementation
func (m MockTrackerAPI) deleteUserJSON(id string) (json []byte, err error) {
	json = []byte(`[true]`)
	return json, nil
}
