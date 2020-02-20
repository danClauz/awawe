package interactor

import (
	"awawe/domain/dto"
	"awawe/domain/model"
	"awawe/usecase/presenter"
	pMock "awawe/usecase/presenter/mock"
	"awawe/usecase/repository"
	rMock "awawe/usecase/repository/mock"
	"context"
	"errors"
	"reflect"
	"testing"
	"time"
)

func postToModel(post *dto.Post) *model.Post {
	return &model.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func Test_postInteractor_FindAll(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name               string
		args               args
		posts              []*dto.Post
		response           interface{}
		postRepositoryResp interface{}
		userRepositoryResp interface{}
		wantErr            bool
	}{
		// TODO: Add test cases.
		{
			name: "find all posts success",
			args: args{
				ctx: context.Background(),
			},
			posts: []*dto.Post{
				{
					ID:        1,
					UserID:    1,
					Title:     "this is title",
					Content:   "this is content",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					User: &dto.User{
						ID:        1,
						Username:  "username",
						FirstName: "first name",
						LastName:  "last name",
						Email:     "user@user.com",
						Password:  "user1234",
					},
				},
				{
					ID:        2,
					UserID:    2,
					Title:     "this is more title",
					Content:   "this is more content",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					User: &dto.User{
						ID:        2,
						Username:  "more username",
						FirstName: "more first name",
						LastName:  "more last name",
						Email:     "moreUser@user.com",
						Password:  "moreUser1234",
					},
				},
			},
			response: []*dto.Post{
				{
					ID:        1,
					UserID:    1,
					Title:     "this is title",
					Content:   "this is content",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					User: &dto.User{
						ID:        1,
						Username:  "username",
						FirstName: "first name",
						LastName:  "last name",
						Email:     "user@user.com",
						Password:  "user1234",
					},
				},
				{
					ID:        2,
					UserID:    2,
					Title:     "this is more title",
					Content:   "this is more content",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					User: &dto.User{
						ID:        2,
						Username:  "more username",
						FirstName: "more first name",
						LastName:  "more last name",
						Email:     "moreUser@user.com",
						Password:  "moreUser1234",
					},
				},
			},
			postRepositoryResp: nil,
			userRepositoryResp: nil,
			wantErr:            false,
		},
		{
			name: "find all posts failed",
			args: args{
				ctx: context.Background(),
			},
			posts:              nil,
			response:           nil,
			postRepositoryResp: errors.New("some error"),
			userRepositoryResp: nil,
			wantErr:            true,
		},
		//{
		//	name: "find all posts failed when get user by id",
		//	args: args{
		//		ctx: context.Background(),
		//	},
		//	posts: []*dto.Post{
		//		{
		//			ID:        1,
		//			UserID:    1,
		//			Title:     "this is title",
		//			Content:   "this is content",
		//			CreatedAt: time.Now(),
		//			UpdatedAt: time.Now(),
		//		},
		//	},
		//	response:           []*dto.Post{},
		//	postRepositoryResp: nil,
		//	userRepositoryResp: errors.New("some error"),
		//	wantErr:            true,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			posts := make([]*model.Post, 0)
			for _, val := range tt.posts {
				posts = append(posts, postToModel(val))
			}

			postRepositoryMock := new(rMock.PostRepositoryMock)
			postRepositoryMock.On("FindAll", tt.args.ctx).Return(posts, tt.postRepositoryResp)

			userRepositoryMock := new(rMock.UserRepositoryMock)
			if tt.postRepositoryResp == nil {
				for key, val := range posts {
					val.User = userToModel(tt.posts[key].User)
					userRepositoryMock.On("GetByID", tt.args.ctx, int(val.UserID)).Return(val.User, tt.userRepositoryResp)
				}
			}

			postPresenterMock := new(pMock.PostPresenterMock)
			if !tt.wantErr {
				postPresenterMock.On("ResponsePosts", posts).Return(tt.response)
			}

			in := NewPostInteractor(postRepositoryMock, userRepositoryMock, postPresenterMock)

			got, err := in.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.response) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.response)
			}

			postRepositoryMock.AssertExpectations(t)
			userRepositoryMock.AssertExpectations(t)
			postPresenterMock.AssertExpectations(t)
		})
	}
}

func Test_postInteractor_Store(t *testing.T) {
	type fields struct {
		postRepository repository.PostRepository
		userRepository repository.UserRepository
		postPresenter  presenter.PostPresenter
	}
	type args struct {
		ctx  context.Context
		post *dto.StorePost
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &postInteractor{
				postRepository: tt.fields.postRepository,
				userRepository: tt.fields.userRepository,
				postPresenter:  tt.fields.postPresenter,
			}
			if err := in.Store(tt.args.ctx, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_postInteractor_Update(t *testing.T) {
	type fields struct {
		postRepository repository.PostRepository
		userRepository repository.UserRepository
		postPresenter  presenter.PostPresenter
	}
	type args struct {
		ctx  context.Context
		post *dto.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := &postInteractor{
				postRepository: tt.fields.postRepository,
				userRepository: tt.fields.userRepository,
				postPresenter:  tt.fields.postPresenter,
			}
			if err := in.Update(tt.args.ctx, tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
