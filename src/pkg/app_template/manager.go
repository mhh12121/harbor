package apptemplate

import (
	"context"

	"github.com/goharbor/harbor/src/pkg/app_template/model"
)

type AppTemplateImpl interface {
	CreateAppTemplate(ctx context.Context, at model.AppTemplate) error
	GetAppTemplate(ctx context.Context, templateID string) (*model.AppTemplate, error)
}
