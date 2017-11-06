package controllers

import "github.com/revel/revel"

func init() {
	revel.InterceptMethod(Auth.AddUser, revel.BEFORE)
	revel.InterceptMethod(Prayers.checkUser, revel.BEFORE)
}
