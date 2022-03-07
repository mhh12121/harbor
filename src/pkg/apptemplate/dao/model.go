package dao

import (
	"github.com/astaxie/beego/orm"
	"github.com/goharbor/harbor/src/lib/q"
	"time"
)

func init() {
	orm.RegisterModel(&AppTemplate{})
}

//app template DB
type AppTemplate struct {
	ID int64 `orm:"pk;column(id)"`
	// GlobalID    string `orm:"column(global_id)"` //generated id
	AppName     string    `orm:"column(app_name)"`
	Desc        string    `orm:"column(desc)"`
	ICON        string    `orm:"column(icon)"`
	Status      int       `orm:"column(status)"`
	ProjectID   int64     `orm:"column(project_id)"`
	ManifestLoc string    `orm:"column(manifest_loc);type(jsonb)"`
	UpdateTime  time.Time `orm:"column(update_time)"`
	CreateTime  time.Time `orm:"column(create_time)"`
	// ExtraAttrs  string    `orm:"column(extra_attrs)"`             //json
	// Annotations string    `orm:"column(annotations);type(jsonb)"` //json
}

func (at *AppTemplate) Table() string {
	return "app_template"
}

// GetDefaultSorts specifies the default sorts
func (at *AppTemplate) GetDefaultSorts() []*q.Sort {
	return []*q.Sort{
		{
			Key:  "PushTime",
			DESC: true,
		},
		{
			Key:  "ID",
			DESC: true,
		},
	}
}
