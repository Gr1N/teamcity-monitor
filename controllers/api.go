package controllers

import (
	"github.com/Gr1N/teamcity-monitor/models"
)

type APIController struct {
	BaseController
}

func (c *APIController) URLMapping() {
	c.Mapping("Builds", c.Builds)
	c.Mapping("BuildsStatus", c.BuildsStatus)
	c.Mapping("BuildsStatistics", c.BuildsStatistics)
}

// @router /builds [get]
func (c *APIController) Builds() {
	builds := models.Builds()

	c.Data["json"] = builds
	c.ServeJson()
}

// @router /builds/status [get]
func (c *APIController) BuildsStatus() {
	buildsStatus := models.BuildsStatus()

	c.Data["json"] = buildsStatus
	c.ServeJson()
}

// @router /builds/statistics [get]
func (c *APIController) BuildsStatistics() {
	// TBD
}
