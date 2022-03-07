package dao

import (
	"context"
	beegoorm "github.com/astaxie/beego/orm"
	"github.com/goharbor/harbor/src/lib/errors"
	"github.com/goharbor/harbor/src/lib/orm"
	"github.com/goharbor/harbor/src/lib/q"
)

type DAO interface {
	Count(ctx context.Context, query *q.Query) (total int64, err error)
	List(ctx context.Context) (appTemplates []AppTemplate, err error)
	Get(ctx context.Context, id int64) (appTemplate *AppTemplate, err error)
	Create(ctx context.Context, appTemplate *AppTemplate) (id int64, globalID string, err error)
	Delete(ctx context.Context, id int64) (err error)
	Update(ctx context.Context, appTemplate *AppTemplate, props ...string) (err error)
}
type dao struct {
}

func New() DAO {
	return &dao{}
}

func (d *dao) Count(ctx context.Context, query *q.Query) (total int64, err error) {
	if query != nil {
		// ignore the page number and size
		query = &q.Query{
			Keywords: query.Keywords,
		}
	}
	qs, err := querySetter(ctx, query)
	if err != nil {
		return 0, err
	}
	return qs.Count()
}
func (d *dao) List(ctx context.Context) (appTemplates []AppTemplate, err error) {
	return
}
func (d *dao) Get(ctx context.Context, id int64) (appTemplate *AppTemplate, err error) {
	return
}
func (d *dao) Create(ctx context.Context, appTemplate *AppTemplate) (id int64, globalID string, err error) {
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return 0, "", err
	}
	id, err = ormer.Insert(appTemplate)
	if err != nil {
		if e := orm.AsConflictError(err, "apptemplate:%s already exists project:%d",
			appTemplate.AppName, appTemplate.ProjectID); e != nil {
			err = e
		}
	}
	return
}
func (d *dao) Delete(ctx context.Context, id int64) (err error) {
	ormer, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}
	n, err := ormer.Delete(&AppTemplate{
		ID: id,
	})
	if err != nil {
		if e := orm.AsForeignKeyError(err,
			"the artifact %d is referenced by other resources", id); e != nil {
			err = e
		}
		return err
	}
	if n == 0 {
		return errors.NotFoundError(nil).WithMessage("artifact %d not found", id)
	}
	return nil
}

func (d *dao) Update(ctx context.Context, appTemplate *AppTemplate, props ...string) (err error) {
	return
}

const (
	both = ""
)

func querySetter(ctx context.Context, query *q.Query) (beegoorm.QuerySeter, error) {
	qs, err := orm.QuerySetter(ctx, &AppTemplate{}, query)
	if err != nil {
		return nil, err
	}
	qs, err = setNormalQuery(qs, query)
	if err != nil {
		return nil, err
	}
	// qs, err = setTagQuery(ctx, qs, query)
	// if err != nil {
	// 	return nil, err
	// }
	// qs, err = setLabelQuery(qs, query)
	// if err != nil {
	// 	return nil, err
	// }
	return qs, nil
}
func setNormalQuery(qs beegoorm.QuerySeter, query *q.Query) (beegoorm.QuerySeter, error) {
	if query == nil || len(query.Keywords) == 0 {
		qs = qs.FilterRaw("id", both)
		return qs, nil
	}
	base, exist := query.Keywords["status"]
	if !exist {
		qs = qs.FilterRaw("id", both)
		return qs, nil
	}
	b, ok := base.(string)
	if !ok || b != "*" {
		return qs, errors.New(nil).WithCode(errors.BadRequestCode).
			WithMessage(`the value of "base" query can only be exact match value with "*"`)
	}
	// the base is specified as "*"
	return qs, nil
}
