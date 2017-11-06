package controllers

import (
	"SocialPrayer/app/routes"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
	Auth
}

func (c App) Index() revel.Result {
	if c.Auth.connected() == nil {
		return c.Redirect(routes.App.Hello())
	}
	loggedIn := true
	return c.Render(loggedIn)
}

func (c App) Hello() revel.Result {
	return c.Render()
}
