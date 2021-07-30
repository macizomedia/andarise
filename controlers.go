package database

import "context"

var collection = GetCollection("subscribers")
var ctx = context.Background()

func Create(user User) error {
	var err error

	_,err = collection.InsertOne(ctx,user)

	if err != nil {
		return err
	}

	return nil
}