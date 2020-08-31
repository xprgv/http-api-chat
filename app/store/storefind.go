package store

import (
	"fmt"

	"github.com/openmind13/http-api-chat/app/model"
)

// FindUserByUsername ...
func (s *SQLStore) FindUserByUsername(username string) (*model.User, error) {
	user := &model.User{}

	if err := s.db.QueryRow(
		"SELECT id, username, created_at FROM users WHERE username = $1;",
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.CreatedAt,
	); err != nil {
		fmt.Println("error")
		return nil, err
	}

	return user, nil

	// row, err := s.db.Query(
	// 	"SELECT id, username, created_at FROM users WHERE username = $1;",
	// 	username,
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// defer row.Close()

	// for row.Next() {
	// 	if err := row.Scan(
	// 		&user.ID,
	// 		&user.Username,
	// 		&user.CreatedAt,
	// 	); err != nil {
	// 		return nil, err
	// 	}

	// 	return user, nil
	// }
}

// GetAllUsers ...
func (s *SQLStore) GetAllUsers() ([]model.User, error) {
	row, err := s.db.Query(
		"SELECT id, username, created_at FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var users []model.User
	// users = append(users, &model.User{})

	i := 0
	for row.Next() {
		users = append(users, model.User{})

		if err := row.Scan(
			&users[i].ID,
			&users[i].Username,
			&users[i].CreatedAt,
		); err != nil {
			return nil, err
		}

		i++
	}

	return users, nil
}
