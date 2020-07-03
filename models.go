// Models for mapping eTracker REST calls to go structs
// Etuser - user informations and login credentials bound to a single account, including assigned role
// Etrole - userrole informations bound to a single account
package trackingapi

// Etuser used for mapping eTracker user information from REST calls to go structs.
// User information and login credentials bound to a single account, including assigned role
// for managing sub-users
type Etuser struct {
	Id       string `json:"id"`        // ID created by eTracker and returned by "create" call
	Name     string `json:"name"`      // Last name
	Fname    string `json:"fname"`     // First name
	Version  string `json:"version"`   // version number
	Subid    string `json:"subid"`     // sub-id created by eTracker as part of the user login "AccountId-Subid" e.g. "1898673-14"
	Rolename string `json:"role"`      // The name for the authorisation profile
	Roletype string `json:"role_type"` // Read and write permissions or just read permissions, value can be: read_only or admin
	Roleid   string `json:"role_id"`   // Identifier of the user profile
	Enable   string `json:"enable"`    // active, value 1 or inactive, value 0
	Login    string `json:"login"`     // count of logins with this sub-account
	Password string `json:"password"`  //  Password
	Email    string `json:"email"`     // Email address
	Language string `json:"language"`  // (Spanisch = es, Englisch = en, Französisch = fr und Deutsch = de)
	Sex      string `json:"sex"`       // 0..male 1..female
}

// Etrole used for mapping eTracker user information from REST calls to go structs.
// Role information bound to a single account for managing authorisation profiles
type Etrole struct {
	Id      string `json:"id"`      // ID of the role, referenced by Etuser
	Name    string `json:"name"`    // The name for the authorisation profile
	Version string `json:"version"` // version number
	Users   string `json:"users"`   // amount of users with this profile
}
