# trackingapi
This package is an API abstraction layer to manage a larger amount of eTracker
accounts (https://www.etracker.com/) in golang.

It is providing functionality for managing Users and Roles using the eTracker
REST API and for testing purpose a mock implementation

## REST API:
eTracker REST API documentation: (https://www.etracker.com/en/docs/integration-setup-2/developer-apis/management-of-authorisation-profiles-an-sub-users/)

## Implementation:
To implement the tracking we have first to create a Tracking struct, that
will hold all created TrackingAccounts
```
tracking := NewTracking()
```

Every TrackingAccount has to be initialised using an eTracker account or
for testing purpose the mock implementation. For this purpose we do have
ETrackerAPI for eTracker and MockTrackerAPI for the mock implementation
```
tracking.AddAccount("TEST", NewTrackingAccount(NewMockTrackerAPI()))
```

TrackingAccount are saved internally as a map[string]TrackingAccount, so you
can list all keys of the map by calling
```
keys:= AllAccountIds()
```

One single account can be accessed by calling with the requested key
```
Account()
```

## Sample:
```
tracking := NewTracking()
tracking.AddAccount("TEST", NewTrackingAccount(NewMockTrackerAPI()))

account, err := tracking.Account("TEST")
if err != nil {
	fmt.Printf("account not found: %v\n", err)
}

ret, err := account.GetRoles()
if err != nil {
	fmt.Printf("roles not found: %v\n", err)
}
fmt.Println(ret)
// Output: [{39 Analytics only Access Rights 1 2} {21 Analyze and Edit Access Rights 1 19}]
```
