package db

func GetAllUsers() ([]User, error) {
	var users []User
	err := DB.Find(&users).Error
	return users, err
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := DB.Find(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(userID uint) (User, error) {
	var user User
	err := DB.Preload("Projects").Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
