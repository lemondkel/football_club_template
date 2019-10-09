package controllers

import (
	"github.com/revel/revel"
)

type Tickets struct {
	*revel.Controller
}

func (c Tickets) Index() revel.Result {
	return c.Render()
}
