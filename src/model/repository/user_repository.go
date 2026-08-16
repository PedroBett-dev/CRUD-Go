package repository

import (
	"context"
	"errors"

	"github.com/PedroBett-dev/CRUD-Go.git/src/model/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(user entity.UserEntity) (string, error) {
	var id string

	err := r.pool.QueryRow(
		context.Background(),
		`INSERT INTO users (name, age, email, password) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Name,
		user.Age,
		user.Email,
		user.Password,
	).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepository) FindByID(id string) (*entity.UserEntity, error) {
	var user entity.UserEntity

	err := r.pool.QueryRow(
		context.Background(),
		`SELECT id, name, age, email, password FROM users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.UserEntity, error) {
	var user entity.UserEntity

	err := r.pool.QueryRow(
		context.Background(),
		`SELECT id, name, age, email, password FROM users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(id string, user entity.UserEntity) error {
	command, err := r.pool.Exec(
		context.Background(),
		`UPDATE users SET name = $1, age = $2, email = $3, password = $4 WHERE id = $5`,
		user.Name,
		user.Age,
		user.Email,
		user.Password,
		id,
	)
	if err != nil {
		return err
	}

	if command.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	command, err := r.pool.Exec(
		context.Background(),
		`DELETE FROM users WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	if command.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func IsNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

func (r *UserRepository) List() ([]entity.UserEntity, error) {
	rows, err := r.pool.Query(
		context.Background(),
		`SELECT id, name, age, email, password FROM users`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.UserEntity

	for rows.Next() {
		var user entity.UserEntity

		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
