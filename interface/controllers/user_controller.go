package controllers

import (
	"awawe/domain/dto"
	"awawe/usecase/interactor"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func StoreUserToRedis(svc interactor.UserInteractor) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := new(dto.User)
		if err := ctx.Bind(user); err != nil {
			return err
		}

		if err := svc.StoreToRedis(ctx.Request().Context(), user); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, nil)
	}
}

func FindAllUsers(svc interactor.UserInteractor) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		users, err := svc.FindAll(ctx.Request().Context())
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, users)
	}
}

func GetUserByID(svc interactor.UserInteractor) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")

		userID, err := strconv.Atoi(id)
		if err != nil {
			return err
		}

		user, err := svc.GetByID(ctx.Request().Context(), userID)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, user)
	}
}
