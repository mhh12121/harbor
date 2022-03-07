package model

import (
	"github.com/goharbor/harbor/src/controller/apptemplate"
	"github.com/goharbor/harbor/src/server/v2.0/models"
)

type AppTemplate struct {
	apptemplate.AppTemplate
}

func (a *AppTemplate) ToSwagger() *models.AppTemplate {
	appt := &models.AppTemplate{
		ID: a.ID,
	}
	return appt
}
