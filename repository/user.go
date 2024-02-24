package repository

import (
	"github.com/heitorlimamorei/go-chat-api/schemas"
)

func GetUsers() ([]*schemas.User, error) {
	var users []*schemas.User

	err := db.Find(&users).Error

	if err != nil {
		logger.ErrorF("Error while getting users: %v", err)
		return nil, err
	}

	return users, nil
}

func GetUserByEmail(email string) (*schemas.User, error) {
	user := &schemas.User{}

	err := db.Where("email =?", email).First(user).Error

	if err != nil {
		logger.ErrorF("Error while getting users: %v", err)
		return nil, err
	}

	return user, nil
}

func GetUserById(id string) (*schemas.User, error) {
	user := schemas.User{}

	err := db.Find(&user, id).Error

	if err != nil {
		logger.ErrorF("Error while getting user: %v", err)
		return nil, err
	}

	return &user, nil
}

func CreateUser(newUser *schemas.User) error {
	err := db.Create(newUser).Error

	if err != nil {
		logger.ErrorF("Error while creating user: %v", err)
		return err
	}
	return nil
}

func DeleteUSer(id string) error {
	user := schemas.User{}

	err := db.First(&user, id).Error

	if err != nil {
		logger.ErrorF("Error while deleting user: %v", err)
		return err
	}

	db.Delete(&user, id)
	return nil
}

func UpdateUser(updatedUser *schemas.User, id string) error {
	user, err := GetUserById(id)

	if err != nil {
		return err
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Age = updatedUser.Age

	db.Save(user)

	return nil
}
