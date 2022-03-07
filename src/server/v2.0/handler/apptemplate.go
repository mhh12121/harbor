package handler

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/goharbor/harbor/src/common/rbac"
	"github.com/goharbor/harbor/src/controller/apptemplate"
	"github.com/goharbor/harbor/src/controller/project"
	pkg_apptemplate "github.com/goharbor/harbor/src/pkg/apptemplate"
	"github.com/goharbor/harbor/src/server/v2.0/handler/model"
	"github.com/goharbor/harbor/src/server/v2.0/models"
	operation "github.com/goharbor/harbor/src/server/v2.0/restapi/operations/apptemplate"
)

func newAppTemplateAPI() *appTemplateAPI {
	return &appTemplateAPI{}

}

type appTemplateAPI struct {
	BaseAPI
	apptCtl apptemplate.Controller
	proCtl  project.Controller
	// repoCtl repository.Controller
	// scanCtl scan.Controller
	// tagCtl  tag.Controller
}

func (a *appTemplateAPI) Prepare(ctx context.Context, operation string, params interface{}) middleware.Responder {
	return nil
}

/* ListAppTemplates List app templates */
func (a *appTemplateAPI) ListAppTemplates(ctx context.Context, params operation.ListAppTemplatesParams) middleware.Responder {
	if err := a.RequireProjectAccess(ctx, params.ProjectName, rbac.ActionList, rbac.ResourceAppTemplate); err != nil {
		return a.SendError(ctx, err)
	}

	// set query
	query, err := a.BuildQuery(ctx, params.Q, params.Sort, params.Page, params.PageSize)
	if err != nil {
		return a.SendError(ctx, err)
	}
	query.Keywords["ProjectName"] = params.ProjectName

	// set option
	// option := option(params.WithTag, params.WithImmutableStatus,
	// 	params.WithLabel, params.WithSignature)
	// option:=
	// get the total count of artifacts
	total, err := a.apptCtl.Count(ctx, query)
	if err != nil {
		return a.SendError(ctx, err)
	}
	// list artifacts according to the query and option
	arts, err := a.apptCtl.List(ctx, query, nil)
	if err != nil {
		return a.SendError(ctx, err)
	}
	//convert to swagger models
	var appTemplates []*models.AppTemplate
	for _, art := range arts {
		a := &model.AppTemplate{}
		a.AppTemplate = *art
		appTemplates = append(appTemplates)
	}
	return operation.NewListAppTemplatesOK().
		WithXTotalCount(total).
		WithLink(a.Links(ctx, params.HTTPRequest.URL, total, query.PageNumber, query.PageSize).String()).
		WithPayload(appTemplates)
}

func (a *appTemplateAPI) AddAppTemplate(ctx context.Context, params operation.AddAppTemplateParams) middleware.Responder {
	if err := a.RequireProjectAccess(ctx, params.ProjectName, rbac.ActionCreate, rbac.ResourceAppTemplate); err != nil {
		return a.SendError(ctx, err)
	}
	appt := params.From

	aptTemplateModel := pkg_apptemplate.AppTemplate{
		ID:       appt.ID,
		GlobalID: appt.GlobalID,
		// AppName:  appt.AppName,
	}
	appTemplate := &apptemplate.AppTemplate{
		AppTemplate: aptTemplateModel,
	}

	err := a.apptCtl.Add(ctx, appTemplate)
	if err != nil {
		return a.SendError(ctx, err)
	}

	return operation.NewAddAppTemplateOK()
}
