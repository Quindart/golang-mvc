package repositories

import (
	"database/sql"
	"example/golang-mvc/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Lấy danh sách tất cả users
func (r *UserRepository) GetAll() ([]models.User, error) {
	query := "select id, fullName, age from users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FullName, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Tạo user mới
func (r *UserRepository) Create(user *models.User) error {
	query := "insert into users (fullName, age) values (?, ?)"
	result, err := r.DB.Exec(query, user.FullName, user.Age)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}