package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) URLMapping() {
	c.Mapping("Index", c.Index)
}

// @router / [get]
func (c *IndexController) Index() {
	c.TplNames = "index.tpl"
}
