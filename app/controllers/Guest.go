package controllers

import (
	"github.com/revel/revel"
)

type GuestUser struct {
	App
}

func (gu GuestUser) Index() revel.Result {
	return gu.Render()
}
