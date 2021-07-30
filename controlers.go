package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var collection = GetCollection("subscribers")
var ctx = context.Background()

func Create(user *User) error {
	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := &User{}
	err := json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		return
	}

	createdUser := CreateOne(newUser)
	var errMessage = createdUser.Error()

	if createdUser != nil {
		fmt.Println(errMessage)
	}
	err = json.NewEncoder(w).Encode(createdUser)
	if err != nil {
		return
	}

}
