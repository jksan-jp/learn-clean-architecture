package database

import (
	"learn-clean-architecture/domain"
	"learn-clean-architecture/infrastructure"
)

type UserRepository struct {
	infrastructure.SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Conn.Exec("INSERT INTO users(first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName)
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

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Conn.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer func() {
		err := row.Close()
		if err != nil {
			return
		}
	}()

	var id int
	var firstName, lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Conn.Query("SELECT id, first_name, last_name FROM users")
	defer func() {
		rows.Close()
		if err != nil {
			return
		}
	}()

	for rows.Next() {
		var id int
		var firstName, lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return
}
