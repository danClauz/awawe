package controllers

import (
	"awawe/domain/dto"
	"awawe/usecase/interactor"
	"github.com/labstack/echo"
	"net/http"
)

func StorePost(svc interactor.PostInteractor) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		post := new(dto.Post)
		if err := ctx.Bind(post); err != nil {
			return err
		}

		if err := svc.Store(ctx.Request().Context(), post); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func UpdatePost(svc interactor.PostInteractor) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		post := new(dto.Post)
		if err := ctx.Bind(post); err != nil {
			return err
		}

		if err := svc.Update(ctx.Request().Context(), post); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func FindAllPosts(svc interactor.PostInteractor) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		posts, err := svc.FindAll(ctx.Request().Context())
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, posts)
	}
}
