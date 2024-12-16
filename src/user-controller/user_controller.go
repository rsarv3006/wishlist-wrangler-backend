package usercontroller

import (
	"net/http"
	coresharedcontroller "wishlist-wrangler-api/src/core-shared/core-shared-controller"
	coresharedrestservice "wishlist-wrangler-api/src/core-shared/core-shared-rest-service"
	usershareddao "wishlist-wrangler-api/src/user-shared/user-shared-dao"
	usersharedservice "wishlist-wrangler-api/src/user-shared/user-shared-service"
)

type UserController struct {
	coresharedcontroller.CrudController[uint, usershareddao.UserEntity, usershareddao.UserEntityDTO, *usersharedservice.UserService]
}

func NewUserController(service *usersharedservice.UserService) *UserController {
	return &UserController{
		CrudController: coresharedcontroller.CrudController[uint, usershareddao.UserEntity, usershareddao.UserEntityDTO, *usersharedservice.UserService]{Service: service},
	}
}

func (c *UserController) RegisterRoutes() []coresharedrestservice.Route {
	return []coresharedrestservice.Route{
		c.RegisterGetAllEntitiesRoute(),
		c.RegisterGetByIdRoute(),
		c.RegisterUpdateDataRoute(),
		c.RegisterCreateUserRoute(),
	}
}

// RegisterGetAllEntitiesRoute
// @Summary Get all users
// @Description Returns all users
// @Tags user
// @Produce json
// @Success 200 {object} []usershareddao.UserEntity
// @Failure 404 {object} map[string]string
// @Failure 501 {object} map[string]string
// @Router /api/users/user/v1 [get]
func (c *UserController) RegisterGetAllEntitiesRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodGet,
		Pattern: "/api/users/user/v1",
		Handler: c.GetAllEntities,
	}
}

// RegisterGetByIdRoute
// @Summary Get user by ID
// @Description Returns user by ID
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} usershareddao.UserEntity
// @Failure 404 {object} map[string]string
// @Router /api/users/user/v1/{id} [get]
func (c *UserController) RegisterGetByIdRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodGet,
		Pattern: "/api/users/user/v1/{id}",
		Handler: c.GetById,
	}
}

// RegisterUpdateDataRoute
// @Summary Update user data
// @Description Updates user data
// @Tags user
// @Produce json
// @Param id path int true "User ID"
// @Param data body usershareddao.UserEntityDTO true "User data"
// @Success 200 {object} usershareddao.UserEntity
// @Failure 404 {object} map[string]string
// @Router /api/users/user/v1/{id} [put]
func (c *UserController) RegisterUpdateDataRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodPut,
		Pattern: "/api/users/user/v1/{id}",
		Handler: c.UpdateData,
	}
}

// RegisterCreateUserRoute
// @Summary Create user
// @Description Creates user
// @Tags user
// @Produce json
// @Param data body usershareddao.UserEntityDTO true "User data"
// @Success 201 {object} usershareddao.UserEntity
// @Failure 404 {object} map[string]string
// @Router /api/users/user/v1 [post]
func (c *UserController) RegisterCreateUserRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodPost,
		Pattern: "/api/users/user/v1",
		Handler: c.CreateData,
	}
}
