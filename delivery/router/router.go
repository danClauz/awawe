package router

import (
	config "awawe/configuration"
	"awawe/delivery/controllers"
	"awawe/infrastructure/datastore"
	"awawe/registry"
	"awawe/usecase/interactor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	user interactor.UserInteractor
	post interactor.PostInteractor
)

func initializeService() {
	db := datastore.NewMySQLDB()

	if config.GetAppConfig().AutoMigrate {
		datastore.MigrateMySQLDatabase(db)
	}

	user = registry.NewUserInteractor(db)
	post = registry.NewPostInteractor(db)
}

func NewRouter() *echo.Echo {
	initializeService()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// User API
	e.POST("/user", controllers.StoreUser(user))
	e.POST("/user-to-redis", controllers.StoreUserToRedis(user))
	e.GET("/users", controllers.FindAllUsers(user))
	e.GET("/user/:id", controllers.GetUserByID(user))

	// Post API
	e.POST("/post", controllers.StorePost(post))
	e.PATCH("/post", controllers.UpdatePost(post))
	e.GET("/posts", controllers.FindAllPosts(post))

	return e
}
