package main

import "testing"

type user struct {
	x int
	y string	
	z string	
	a string	
}

var getSession = []GetUser{
	{1, "zaina", "zaina", "zaina123"}
}


func TestgetUser(t *testing.T) {
	for _, test := range getSession {
		result, err := getUser
		if(result !=getSession){
			t.Fatalf("Expected Result Not given")
		}
	}
}


