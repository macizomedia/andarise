package database

import (
	"testing"
	"time"
)

func TestCreateOne(t *testing.T) {
	user := User{
		Name: "Blas",
		Email: "mail@server.com",
		Phone: "017630725623",
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