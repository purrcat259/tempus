package db

func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}
