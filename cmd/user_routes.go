package cmd

import "github.com/shawgichan/tourist/repo/user"

func (server *RouteServer) UserRoutes() {
	users := user.NewPlaceServer(server.Store)

	d := server.Router.Group("api/auth")
	{
		d.POST("/registerUser", users.RegisterUser)
	}
}
