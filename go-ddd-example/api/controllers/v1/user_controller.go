package controllers_v1

import (
	context2 "context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-ddd-example/application/users"
	domainUsers "go-ddd-example/domain/users"
	"net/http"
)

const _prefix = "/users"

func CreateGuestUser(group *echo.Group, userService users.UserService) {
	path := fmt.Sprintf("%s/CreateGuestUser", _prefix)
	group.POST(path, func(context echo.Context) error {

		var (
			user *domainUsers.User
			err  error
		)

		if user, err = userService.AddNewGuestUser(context2.Background()); err != nil {
			return err
		}

		return context.JSON(http.StatusCreated, user)
	})
}

func GetUserByObjectId(group *echo.Group, userService users.UserService) {
	path := fmt.Sprintf("%s/id/:id", _prefix)
	group.GET(path, func(context echo.Context) error {

		var (
			user *domainUsers.User
			err  error
		)

		id := context.Param("id")
		if user, err = userService.GetUserById(context2.Background(), id); err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}

		return context.JSON(http.StatusCreated, user)
	})
}
