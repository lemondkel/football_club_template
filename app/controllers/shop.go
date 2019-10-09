package controllers

import (
	"github.com/revel/revel"
)

type Shop struct {
	*revel.Controller
}

func (c Shop) Index() revel.Result {
	return c.Render()
}
