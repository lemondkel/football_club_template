package controllers

import (
	"github.com/revel/revel"
)

type Club struct {
	*revel.Controller
}

func (c Club) Index() revel.Result {
	return c.Render()
}
