package controller

import (
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var userRepo *repository.UserRepository

func InitRepository(pool *pgxpool.Pool) {
	userRepo = repository.NewUserRepository(pool)
}