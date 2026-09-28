package db

import "example/golang-mvc/models"

type UserRepository interface {
	findAll() []models.User
}
