// trackingapi project doc.go

// This package is an API abstraction layer to manage a larger amount of eTracker
// accounts https://www.etracker.com/ in golang.
//
// It is providing functionality for managing Users and Roles using the eTracker
// REST API and for testing purpose a mock implementation
//
// eTracker REST API documentation: https://www.etracker.com/en/docs/integration-setup-2/developer-apis/management-of-authorisation-profiles-an-sub-users/
//
// Implementation:
// To implement the tracking we have first to create a Tracking struct, that
// will hold all created TrackingAccounts
//  tracking := NewTracking()
//
// Every TrackingAccount has to be initialised using an eTracker account or
// for testing purpose the mock implementation. For this purpose we do have
// ETrackerAPI for eTracker and MockTrackerAPI for the mock implementation
// 	tracking.AddAccount("TEST", NewTrackingAccount(NewMockTrackerAPI()))
//
// TrackingAccount are saved internally as a map[string]TrackingAccount, so you
// can list all keys of the map by calling
//  keys:= AllAccountIds()
//
// One single account can be accessed by calling with the requested key
//  account := Account("Account ID")
package trackingapi
