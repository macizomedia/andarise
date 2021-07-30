package main

import (
	"testing"
	"time"
)

func TestCreateOne(t *testing.T) {
	user := &User{
		Name:      "Jhon",
		Email:     "mailing@server.com",
		Phone:     "01763078827",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := Create(user)

	if err != nil {
		t.Error("fail")
		t.Fail()
	} else {
		t.Log("user created")
	}

}
