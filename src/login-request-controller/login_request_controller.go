package usercontroller

import (
	"net/http"
	coresharedcontroller "wishlist-wrangler-api/src/core-shared/core-shared-controller"
	coresharedrestservice "wishlist-wrangler-api/src/core-shared/core-shared-rest-service"
	loginrequestshareddao "wishlist-wrangler-api/src/login-request-shared/login-request-shared-dao"
	loginrequestsharedservice "wishlist-wrangler-api/src/login-request-shared/login-request-shared-service"
)

type LoginController struct {
	coresharedcontroller.CrudController[uint, loginrequestshareddao.LoginRequestEntity, loginrequestshareddao.LoginRequestEntityDTO, *loginrequestsharedservice.LoginRequestService]
}

func NewLoginController(service *loginrequestsharedservice.LoginRequestService) *LoginController {
	return &LoginController{
		CrudController: coresharedcontroller.CrudController[uint, loginrequestshareddao.LoginRequestEntity, loginrequestshareddao.LoginRequestEntityDTO, *loginrequestsharedservice.LoginRequestService]{Service: service},
	}
}

func (c *LoginController) RegisterRoutes() []coresharedrestservice.Route {
	return []coresharedrestservice.Route{
		c.RegisterLoginRoute(),
		c.RegisterSignupRoute(),
		c.RegisterTokenRoute(),
	}
}

// RegisterLoginRoute
// @Summary User login
// @Description Authenticates user and returns a token
// @Tags login-request
// @Produce json
// @Param data body usershareddao.UserLoginDTO true "Login data"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/login/v1 [post]
func (c *LoginController) RegisterLoginRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodPost,
		Pattern: "/api/login/v1",
		Handler: coresharedcontroller.NotImplementedHandler,
	}
}

// RegisterSignupRoute
// @Summary User signup
// @Description Registers a new user
// @Tags login-request
// @Produce json
// @Param data body usershareddao.UserSignupDTO true "Signup data"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/signup/v1 [post]
func (c *LoginController) RegisterSignupRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodPost,
		Pattern: "/api/signup/v1",
		Handler: coresharedcontroller.NotImplementedHandler,
	}
}

// RegisterTokenRoute
// @Summary Refresh token
// @Description Refreshes user token
// @Tags login-request
// @Produce json
// @Param data body usershareddao.UserTokenDTO true "Token data"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/token/v1 [post]
func (c *LoginController) RegisterTokenRoute() coresharedrestservice.Route {
	return coresharedrestservice.Route{
		Method:  http.MethodPost,
		Pattern: "/api/token/v1",
		Handler: coresharedcontroller.NotImplementedHandler,
	}
}
