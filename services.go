package database

func CreateOne(user User) error {
	err := Create(user)

	if err != nil {
		return err
	}
	return nil
}