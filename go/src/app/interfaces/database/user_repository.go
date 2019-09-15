package database

import (
	"app/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute("INSERT INTO users (first_name, last_name) VALUES (?, ?)", u.FirstName, u.LastName)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(id int) (user domain.User, err error) {
	row, err := repo.Query(
		"SELECT first_name, last_name FROM users WHERE id = ?", id,
	)
	defer row.Close()
	if err != nil {
		return
	}
	row.Next()
	var (
		firstName string
		lastName  string
	)
	if err = row.Scan(&firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return
}

func (repo *UserRepository) GetAll() (users domain.Users, err error) {
	rows, err := repo.Query(
		"SELECT id, first_name, last_name FROM users",
	)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		user := domain.User{}
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
