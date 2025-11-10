package repository

import (
	"context"
	"time"
	db "userdata-api/db/sqlc"
)

type UserRepository struct {
	queries *db.Queries
}

func NewUserRepository(conn db.DBTX) *UserRepository {
	return &UserRepository{queries: db.New(conn)}
}

func (repo *UserRepository) CreateUser(ctx context.Context, name string, Dob time.Time) (int32, error) {
	return repo.queries.CreateUser(
		ctx, db.CreateUserParams{
			Name: name,
			Dob:  Dob,
		},
	)
}

func (repo *UserRepository) UpdateUser(ctx context.Context, id int32, name string, Dob time.Time) (*db.User, error) {
	user, err := repo.queries.UpdateUser(
		ctx, db.UpdateUserParams{
			ID:   id,
			Name: name,
			Dob:  Dob,
		},
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return repo.queries.DeleteUser(ctx, id)
}

func (repo *UserRepository) GetUserByID(ctx context.Context, id int32) (*db.User, error) {
	user, err := repo.queries.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) ListUsers(ctx context.Context, limit int32, offset int32) ([]db.User, error) {
	return repo.queries.ListUsers(
		ctx, db.ListUsersParams{
			Limit:  limit,
			Offset: offset,
		},
	)
}
