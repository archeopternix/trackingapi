package trackingapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Tracking holds a map of 0..n TrackingAccounts
type Tracking struct {
	accounts map[string]TrackingAccount
}

// NewTracking creates a new instance of Tracking and initializes internal data structures
func NewTracking() Tracking {
	var t Tracking
	t.accounts = make(map[string]TrackingAccount)
	return t
}

// AddAccount adds a new TrackingAccount to Tracking
func (t *Tracking) AddAccount(accountid string, account TrackingAccount) error {
	_, exists := t.accounts[accountid]
	if exists {
		return fmt.Errorf("cannot add account with id: %v it already exists", accountid)
	}
	t.accounts[accountid] = account
	return nil
}

// AllAccountIds gets all ID's of TrackingAccounts
func (t Tracking) AllAccountIds() (accounts []string) {
	for key := range t.accounts {
		accounts = append(accounts, key)
	}
	return accounts
}

// Account returns the pointer to one TrackingAccount by asking for the key
func (t Tracking) Account(key string) (*TrackingAccount, error) {
	_, exists := t.accounts[key]
	if !exists {
		return nil, fmt.Errorf("account with id: %v does not exist", key)
	}
	acc := t.accounts[key]
	return &acc, nil

}

// TrackingAPI is a generic interface to tracking API's with methods, functions provides
// raw JSON results and can be mocked by 'MockTrackerAPI' or implemented by 'ETrackerAPI'
type TrackingAPI interface {
	getUserByIDJSON(id string) (json []byte, err error)
	getUsersJSON() (json []byte, err error)
	getRolesJSON() (json []byte, err error)
	createUserJSON(data *url.Values) (json []byte, err error)
	deleteUserJSON(id string) (json []byte, err error)
}

// TrackingAccount is the implementation of the tracking API.
// it has to be initialised with an concrete tracking implementation
// (eTrackerAPI or MockAPI for testing purpose)
type TrackingAccount struct {
	TrackingAPI
}

// NewTrackingAccount creates a new TrackingAccount and has to be initialised
// with an concrete tracking implementation (eTrackerAPI or MockAPI for testing purpose)
func NewTrackingAccount(t TrackingAPI) TrackingAccount {

	account := TrackingAccount{TrackingAPI: t}
	return account
}

// GetUsers returns all user records for a specific eTracking account
func (a TrackingAccount) GetUsers() (u []Etuser, err error) {
	t0 := time.Now()

	// calls implementation of IJsonApi
	userjson, err := a.getUsersJSON()
	if err != nil {
		return nil, err
	}

	var arr [][]Etuser

	if err = json.Unmarshal(userjson, &arr); err != nil {
		return nil, fmt.Errorf("GetUsers: %w", err)
	}

	t1 := time.Now()
	log.Printf("Read %v user records from eTracker: %v", strconv.Itoa(len(arr[0])), t1.Sub(t0))
	return arr[0], nil
}

// GetUserByID returns the user record with the unique 'id' for a specific eTracker account
func (a TrackingAccount) GetUserByID(id string) (u Etuser, err error) {
	t0 := time.Now()
	var user Etuser

	// calls implementation of IJsonApi
	userjson, err := a.getUserByIDJSON(id)
	if err != nil {
		return user, err
	}

	var arr []Etuser

	//Convert JSON into golang struct
	if err = json.Unmarshal(userjson, &arr); err != nil {
		return user, fmt.Errorf("GetUserbyId: %w", err)
	}

	t1 := time.Now()
	user = arr[0]
	log.Printf("Read user record %v from eTracker: %v", id, t1.Sub(t0))
	return user, nil
}

// GetRoles returns all roles for a specific eTracker account
func (a TrackingAccount) GetRoles() (r []Etrole, err error) {
	t0 := time.Now()

	rolejson, err := a.getRolesJSON()
	if err != nil {
		return nil, err
	}

	var arr [][]Etrole

	//Convert JSON into golang struct
	if err = json.Unmarshal(rolejson, &arr); err != nil {
		return nil, err
	}

	t1 := time.Now()
	log.Printf("Read %v role records from eTracker: %v", strconv.Itoa(len(arr[0])), t1.Sub(t0))
	return arr[0], nil
}

// CreateUser creates a new user record within one eTracker account
func (a TrackingAccount) CreateUser(user Etuser) (id int, err error) {
	var body []byte
	t0 := time.Now()

	data := url.Values{}
	data.Set("enable", user.Enable)
	data.Set("sex", user.Sex)
	data.Set("fname", user.Fname)
	data.Set("email", user.Email)
	data.Set("name", user.Name)
	data.Set("role_id", user.Roleid)
	data.Set("language", user.Language)
	data.Set("pass", user.Password)

	body, err = a.createUserJSON(&data)
	if err != nil {
		return -1, fmt.Errorf("DeleteEtrackerUser: %v", err)
	}

	// Read created user id out of the json response
	dec := json.NewDecoder(strings.NewReader(string(body)))

	respstruct := []struct {
		Id int
	}{}

	err = dec.Decode(&respstruct)
	if err != nil {
		return -1, fmt.Errorf("CreateUser: %s \n%v", body, err)
	}
	id = respstruct[0].Id // user ID created by eTracker

	t1 := time.Now()

	log.Printf("Create user: %s,%s got new id: %v, %v\n", user.Name, user.Fname, id, t1.Sub(t0))
	return id, nil
}

// DeleteUser deletes a user with a given user id from eTracker
func (a TrackingAccount) DeleteUser(userid string) (err error) {
	var body []byte
	t0 := time.Now()

	body, err = a.deleteUserJSON(userid)
	if err != nil {
		return fmt.Errorf("DeleteUser: %v, %v", err, body)
	}

	t1 := time.Now()

	log.Printf("user with id: %v, %v deleted\n", userid, t1.Sub(t0))
	return nil
}
