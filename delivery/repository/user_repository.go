package repository

import (
	"awawe/domain/model"
	iRedis "awawe/usecase/infrastructure/redis"
	"awawe/usecase/repository"
	"context"
	"database/sql"
)

type userRepository struct {
	db    *sql.DB
	redis iRedis.SDK
}

func NewUserRepository(db *sql.DB, redis iRedis.SDK) repository.UserRepository {
	return &userRepository{
		db:    db,
		redis: redis,
	}
}

func (r *userRepository) Store(ctx context.Context, user *model.User) error {
	if _, err := r.db.ExecContext(ctx, `
		INSERT INTO users (username, first_name, last_name, email, password)
		VALUES (?, ?, ?, ?, ?)
	`, user.Username, user.FirstName, user.LastName, user.Email, user.Password);
		err != nil {
		return err
	}

	return nil
}

func (r *userRepository) StoreToRedis(ctx context.Context, user *model.User) error {
	return r.redis.Set(user.TableName(), user, 0)
}

func (r *userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT * FROM users
	`)

	if err != nil {
		return nil, err
	}

	response, err := r.rowsToModel(rows)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *userRepository) GetByID(ctx context.Context, ID int) (*model.User, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT * FROM users
		WHERE id = ?
	`, ID)

	response, err := r.rowToModel(row)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *userRepository) rowToModel(row *sql.Row) (*model.User, error) {
	user := new(model.User)
	if err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) rowsToModel(rows *sql.Rows) ([]*model.User, error) {
	users := make([]*model.User, 0)

	for rows.Next() {
		user := new(model.User)
		if err := rows.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
