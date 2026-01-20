package repository

import (
	"database/sql"
	"time"
	"videocall/internal/app/users/schema"
)

type userRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) UsersRepositoryInterface {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUsers(user *schema.UserRequest) error {

	stmt := "insert into users (first_name, last_name, username, phone, created_at, created_at_unix, updated_at, updated_at_unix) values ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := repo.db.Exec(stmt,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Phone,
		time.Now(),
		time.Now().UTC().UnixMilli(),
		time.Now(),
		time.Now().UTC().UnixMilli(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) UpdateUsers(id int, user *schema.UserRequest) error {
	stmt := "UPDATE users SET username = $1, phone = $2 WHERE id = $3"

	_, err := repo.db.Exec(stmt,
		user.Username,
		user.Phone,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) DeleteUsers(id int) error {
	stmt := "DELETE FROM users WHERE id = $1"

	_, err := repo.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) GetUserById(id int) (*schema.Users, error) {
	stmt := "SELECT * FROM users WHERE id = $1"

	var user schema.Users

	err := repo.db.QueryRow(stmt, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Phone, &user.Created_at, &user.Created_at_unix, &user.Updated_at, &user.Updated_at_unix)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetAllUser() ([]*schema.Users, error) {
	stmt := "SELECT * FROM users"

	row, err := repo.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var user_data []*schema.Users

	for row.Next() {
		var user schema.Users
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Username, &user.Phone, &user.Created_at, &user.Created_at_unix, &user.Updated_at, &user.Updated_at_unix)
		if err != nil {
			return nil, err
		}

		user_data = append(user_data, &user)
	}

	return user_data, nil
}
