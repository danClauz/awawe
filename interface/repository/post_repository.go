package repository

import (
	"awawe/domain/model"
	"awawe/usecase/repository"
	"context"
	"database/sql"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Store(ctx context.Context, post *model.Post) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO posts (user_id, title, content)
		VALUES (?, ?, ?)
	`, post.UserID, post.Title, post.Content)

	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) Update(ctx context.Context, post *model.Post) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE posts
		SET title = ?, content = ?
		WHERE id = ?
	`, post.Title, post.Content, post.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) FindAll(ctx context.Context) ([]*model.Post, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT * FROM posts
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

func (r *postRepository) rowToModel(row *sql.Row) (*model.Post, error) {
	post := new(model.Post)
	if err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepository) rowsToModel(rows *sql.Rows) ([]*model.Post, error) {
	posts := make([]*model.Post, 0)

	for rows.Next() {
		post := new(model.Post)
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
