package repository

import (
	"awawe/domain/model"
	"awawe/usecase/infrastructure/redis/mock"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"regexp"
	"testing"
	"time"
)

func Test_userRepository_Store(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "store user to database success",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					Username:  "danclauz",
					FirstName: "danny",
					LastName:  "ferian",
					Email:     "icanfly654@gmail.com",
					Password:  "1q2w3e4r5t6y",
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

			user := tt.args.user

			sql.ExpectExec("INSERT INTO users").
				WithArgs(user.Username, user.FirstName, user.LastName, user.Email, user.Password).
				WillReturnResult(sqlmock.NewResult(1, 1))

			redisSdkMock := new(mock.RedisSdkMock)

			r := NewUserRepository(db, redisSdkMock)

			if err := r.Store(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}

			err = sql.ExpectationsWereMet()
			assert.NoError(err)

			redisSdkMock.AssertExpectations(t)
		})
	}
}

func Test_userRepository_StoreToRedis(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "store user to redis success",
			args: args{
				ctx: context.Background(),
				user: &model.User{
					ID:        1,
					Username:  "DarrenCOD",
					FirstName: "Darren",
					LastName:  "Cavel",
					Email:     "darren.cavel@ovo.id",
					Password:  "12345678",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, _, err := sqlmock.New()
			assert.NoError(err)
			defer db.Close()

			redisSdkMock := new(mock.RedisSdkMock)

			redisSdkMock.On("Set", tt.args.user.TableName(), tt.args.user, 0*time.Second).
				Return(nil)

			r := NewUserRepository(db, redisSdkMock)
			if err := r.StoreToRedis(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("StoreToRedis() error = %v, wantErr %v", err, tt.wantErr)
			}

			redisSdkMock.AssertExpectations(t)
		})
	}
}

func Test_userRepository_FindAll(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "find all users success",
			args: args{
				ctx: context.Background(),
			},
			want: []*model.User{
				{
					ID:        1,
					Username:  "DanClauz",
					FirstName: "danny",
					LastName:  "ferian",
					Email:     "danny.ferian@ovo.id",
					Password:  "1q2w3e4r5t",
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
				"username",
				"first_name",
				"last_name",
				"email",
				"password",
				"created_at",
				"updated_at",
			}

			rows := sqlmock.NewRows(colName)
			for _, val := range tt.want {
				rows.AddRow(val.ID, val.Username, val.FirstName, val.LastName, val.Email, val.Password, val.CreatedAt, val.UpdatedAt)
			}

			sql.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users")).
				WillReturnRows(rows)

			redisSdkMock := new(mock.RedisSdkMock)

			r := NewUserRepository(db, redisSdkMock)

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

			redisSdkMock.AssertExpectations(t)
		})
	}
}

func Test_userRepository_GetByID(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get user by user id success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: &model.User{
				ID:        1,
				Username:  "DanClauz",
				FirstName: "danny",
				LastName:  "ferian",
				Email:     "danny.ferian@ovo.id",
				Password:  "1q2w3e4r5t",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
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
				"username",
				"first_name",
				"last_name",
				"email",
				"password",
				"created_at",
				"updated_at",
			}

			user := tt.want

			sql.ExpectQuery(regexp.QuoteMeta("SELECT * FROM users")).
				WillReturnRows(sqlmock.NewRows(colName).
					AddRow(user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.CreatedAt, user.UpdatedAt))

			redisSdkMock := new(mock.RedisSdkMock)

			r := NewUserRepository(db, redisSdkMock)

			got, err := r.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}

			err = sql.ExpectationsWereMet()
			assert.NoError(err)

			redisSdkMock.AssertExpectations(t)
		})
	}
}
