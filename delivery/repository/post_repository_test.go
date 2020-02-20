package repository

import (
	"awawe/domain/model"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func Test_postRepository_Store(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx  context.Context
		post *model.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "store post success",
			args: args{
				ctx: context.Background(),
				post: &model.Post{
					UserID:  1,
					Title:   "this is title",
					Content: "this is content",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sql, err := sqlmock.New()
			assert.NoError(err)
			defer db.Close()

			post := tt.args.post

			sql.ExpectExec("INSERT INTO posts").
				WithArgs(post.UserID, post.Title, post.Content).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := NewPostRepository(db)

			if err := r.Store(tt.args.ctx, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}

			err = sql.ExpectationsWereMet()
			assert.NoError(err)
		})
	}
}

func Test_postRepository_Update(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx  context.Context
		post *model.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "update post success",
			args: args{
				ctx: context.Background(),
				post: &model.Post{
					ID:      1,
					Title:   "this is updated title",
					Content: "this is content",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sql, err := sqlmock.New()
			assert.NoError(err)
			defer db.Close()

			post := tt.args.post

			sql.ExpectExec("UPDATE posts").
				WithArgs(post.Title, post.Content, post.ID).
				WillReturnResult(sqlmock.NewResult(1, 1))

			r := NewPostRepository(db)

			if err := r.Update(tt.args.ctx, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}

			err = sql.ExpectationsWereMet()
			assert.NoError(err)
		})
	}
}

func Test_postRepository_FindAll(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Post
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "find all post success",
			args: args{
				ctx: context.Background(),
			},
			want: []*model.Post{
				{
					ID:        1,
					UserID:    1,
					Title:     "this is title",
					Content:   "this is content",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sql, err := sqlmock.New()
			assert.NoError(err)
			defer db.Close()

			colName := []string{
				"id",
				"user_id",
				"title",
				"content",
				"created_at",
				"updated_at",
			}

			rows := sqlmock.NewRows(colName)
			for _, val := range tt.want {
				rows.AddRow(val.ID, val.UserID, val.Title, val.Content, val.CreatedAt, val.UpdatedAt)
			}

			sql.ExpectQuery(regexp.QuoteMeta("SELECT * FROM posts")).
				WillReturnRows(rows)

			r := NewPostRepository(db)

			got, err := r.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}

			err = sql.ExpectationsWereMet()
			assert.NoError(err)
		})
	}
}
