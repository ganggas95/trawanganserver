package controllers

import "github.com/revel/revel"

func init() {

	revel.InterceptMethod(App.SetUser, revel.BEFORE)
	revel.InterceptMethod(Persons.checkUser, revel.BEFORE)
}
