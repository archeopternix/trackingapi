// Example for package eTracker
// GetUsers
package main

import (
	"fmt"
)

func Example() {
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
}

func ExampleTrackingAccount_CreateUser() {
	api := NewTrackingAccount(NewMockTrackerAPI())

	u := Etuser{Id: "72", Enable: "1", Sex: "1", Fname: "XX", Email: "andy@test.de", Name: "XX", Roleid: "21", Language: "de", Password: "12344"}
	ret, err := api.CreateUser(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
	// Output: 78
}

func ExampleTrackingAccount_DeleteUser() {
	api := NewTrackingAccount(NewMockTrackerAPI())

	err := api.DeleteUser("78")
	if err != nil {
		panic(err)
	}

	fmt.Println("deleted")
	// Output: deleted
}

func ExampleTrackingAccount_GetUsers() {
	api := NewTrackingAccount(NewMockTrackerAPI())

	ret, err := api.GetUsers()
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: [{22 Anja Musterfrau 1 14 Analyze and Edit Access Rights read_only 21 1 8    } {25 Bluemchen Benjamin 1 17 Analyze and Edit Access Rights read_only 21 1 8    }]
}

func ExampleTrackingAccount_GetRoles() {
	api := NewTrackingAccount(NewMockTrackerAPI())

	ret, err := api.GetRoles()
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: [{39 Analytics only Access Rights 1 2} {21 Analyze and Edit Access Rights 1 19}]
}

func ExampleTrackingAccount_GetUserByID() {
	api := NewTrackingAccount(NewMockTrackerAPI())

	ret, err := api.GetUserByID("22")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
	// Output: {22 Anja Musterfrau 1 14   21 1 8  a.m@osram.com en 1}
}
