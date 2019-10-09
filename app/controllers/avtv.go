package controllers

import (
	"fct/app/vo"
	"github.com/revel/revel"
)

type AVTV struct {
	*revel.Controller
}

func (c AVTV) Index() revel.Result {
	return c.Render()
}

func (c AVTV) MessageTest() revel.Result {
	m := vo.AjaxResult{Name: "김규빈", Body: "여소받으러 갔네", Time: 1294706395881547000}
	return c.RenderJSON(m)
}
