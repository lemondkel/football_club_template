package controllers

import (
	"github.com/revel/revel"
)

type Matches struct {
	*revel.Controller
}

func (c Matches) Index() revel.Result {
	return c.Render()
}
