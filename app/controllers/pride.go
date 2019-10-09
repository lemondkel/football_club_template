package controllers

import (
	"github.com/revel/revel"
)

type Pride struct {
	*revel.Controller
}

func (c Pride) Index() revel.Result {
	return c.Render()
}
