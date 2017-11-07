package controllers

import (
	"SocialPrayer/app/routes"

	"github.com/revel/revel"
)

type Prayers struct {
	*revel.Controller
	Auth
}

func (c Prayers) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Hello())
	}
	return nil
}

func (c Prayers) Home() revel.Result {
	loggedIn := true
	user := c.connected()
	return c.Render(loggedIn, user)
}
