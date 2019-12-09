package controllers

import (
	"fct/app/vo"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	pageInfo := vo.PageInfo{Title: "asd"}
	return c.Render(pageInfo)
}
