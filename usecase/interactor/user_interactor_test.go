package interactor

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	pMock "awawe/usecase/presenter/mock"
	rMock "awawe/usecase/repository/mock"
	"context"
	"errors"
	"reflect"
	"testing"
)

func userToModel(user *dto.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
}

func Test_userInteractor_Store(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *dto.User
	}
	tests := []struct {
		name               string
		args               args
		repositoryResponse interface{}
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "store user success",
			args: args{
				ctx: context.Background(),
				user: &dto.User{
					Username:  "danclauz",
					FirstName: "danny",
					LastName:  "ferian",
					Email:     "icanfly654@gmail.com",
					Password:  "1q2w3e4r",
				},
			},
			repositoryResponse: nil,
			wantErr:            false,
		},
		{
			name: "store user failed",
			args: args{
				ctx:  context.Background(),
				user: &dto.User{},
			},
			repositoryResponse: errors.New("some error"),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := userToModel(tt.args.user)

			userPresenterMock := new(pMock.UserPresenterMock)
			userPresenterMock.On("RequestToModel", tt.args.user).Return(user)

			userRepositoryMock := new(rMock.UserRepositoryMock)
			userRepositoryMock.On("Store", tt.args.ctx, user).Return(tt.repositoryResponse)

			in := NewUserInteractor(userRepositoryMock, userPresenterMock)

			if err := in.Store(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}

			userRepositoryMock.AssertExpectations(t)
			userPresenterMock.AssertExpectations(t)
		})
	}
}

func Test_userInteractor_StoreToRedis(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *dto.User
	}
	tests := []struct {
		name               string
		args               args
		repositoryResponse interface{}
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "store user to redis success",
			args: args{
				ctx: context.Background(),
				user: &dto.User{
					Username:  "danclauz",
					FirstName: "danny",
					LastName:  "ferian",
					Email:     "icanfly654@gmail.com",
					Password:  "1q2w3e4r",
				},
			},
			repositoryResponse: nil,
			wantErr:            false,
		},
		{
			name: "store user failed",
			args: args{
				ctx:  context.Background(),
				user: &dto.User{},
			},
			repositoryResponse: errors.New("some error"),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := userToModel(tt.args.user)

			userPresenterMock := new(pMock.UserPresenterMock)
			userPresenterMock.On("RequestToModel", tt.args.user).Return(user)

			userRepositoryMock := new(rMock.UserRepositoryMock)
			userRepositoryMock.On("StoreToRedis", tt.args.ctx, user).Return(tt.repositoryResponse)

			in := NewUserInteractor(userRepositoryMock, userPresenterMock)

			if err := in.StoreToRedis(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("StoreToRedis() error = %v, wantErr %v", err, tt.wantErr)
			}

			userRepositoryMock.AssertExpectations(t)
			userPresenterMock.AssertExpectations(t)
		})
	}
}

func Test_userInteractor_FindAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name               string
		args               args
		want               []*dto.User
		repositoryResponse interface{}
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "find all users success",
			args: args{
				ctx: context.Background(),
			},
			want: []*dto.User{
				{
					ID:        1,
					Username:  "danclauz",
					FirstName: "danny",
					LastName:  "ferian",
					Email:     "email@email.com",
				},
				{
					ID:        2,
					Username:  "willsmith",
					FirstName: "willy",
					LastName:  "setiawan",
					Email:     "email@email.com",
				},
			},
			repositoryResponse: nil,
			wantErr:            false,
		},
		{
			name: "find all users failed",
			args: args{
				ctx: context.Background(),
			},
			want:               nil,
			repositoryResponse: errors.New("some error"),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := make([]*model.User, 0)
			for _, val := range tt.want {
				users = append(users, userToModel(val))
			}

			userPresenterMock := new(pMock.UserPresenterMock)
			if !tt.wantErr {
				userPresenterMock.On("ResponseUsers", users).Return(tt.want)
			}

			userRepositoryMock := new(rMock.UserRepositoryMock)
			userRepositoryMock.On("FindAll", tt.args.ctx).Return(users, tt.repositoryResponse)

			in := NewUserInteractor(userRepositoryMock, userPresenterMock)

			got, err := in.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}

			userPresenterMock.AssertExpectations(t)
			userRepositoryMock.AssertExpectations(t)
		})
	}
}

func Test_userInteractor_GetByID(t *testing.T) {
	type args struct {
		ctx context.Context
		ID  int
	}
	tests := []struct {
		name               string
		args               args
		want               *dto.User
		repositoryResponse interface{}
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "get user by id success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want: &dto.User{
				ID:        1,
				Username:  "danclauz",
				FirstName: "danny",
				LastName:  "ferian",
				Email:     "email@email.com",
				Password:  "1q2w3e4r5t6y",
			},
			repositoryResponse: nil,
			wantErr:            false,
		},
		{
			name: "get user by id success",
			args: args{
				ctx: context.Background(),
				ID:  1,
			},
			want:               nil,
			repositoryResponse: errors.New("some error"),
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := new(model.User)
			if tt.want != nil {
				user = userToModel(tt.want)
			} else {
				user = nil
			}

			userPresenterMock := new(pMock.UserPresenterMock)
			if !tt.wantErr {
				userPresenterMock.On("ResponseUser", user).Return(tt.want)
			}

			userRepositoryMock := new(rMock.UserRepositoryMock)
			userRepositoryMock.On("GetByID", tt.args.ctx, tt.args.ID).Return(user, tt.repositoryResponse)

			in := NewUserInteractor(userRepositoryMock, userPresenterMock)

			got, err := in.GetByID(tt.args.ctx, tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}

			userPresenterMock.AssertExpectations(t)
			userRepositoryMock.AssertExpectations(t)
		})
	}
}
