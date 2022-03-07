package apptemplate

import (
	"context"
	"github.com/goharbor/harbor/src/lib/q"
	"github.com/goharbor/harbor/src/pkg/apptemplate/dao"
)

var (
	Mgr = NewManager()
)

type Manager interface {
	// Count returns the total count of apptemplates according to the query.
	Count(ctx context.Context, query *q.Query) (total int64, err error)
	// List apptemplate according to the query. The artifacts that referenced by others and
	// without tags are not returned
	List(ctx context.Context, query *q.Query) (appTemplates []*AppTemplate, err error)
	// Get the apptemplate specified by the ID
	Get(ctx context.Context, id int64) (appTemplate *AppTemplate, err error)
	GetByGlobalID(ctx context.Context, globalid string) (appTemplate *AppTemplate, err error)
	Add(ctx context.Context, appTemplate *AppTemplate) (id int64, globalID string, err error)
	// Create the apptemplate. If the artifact is an index, make sure all the artifacts it references
	// already exist
	Create(ctx context.Context, appTemplate *AppTemplate) (id int64, err error)
	// Delete just deletes the artifact record. The underlying data of registry will be
	// removed during garbage collection
	Delete(ctx context.Context, id int64) (err error)
	//LogicDelete
	DeleteLogic(ctx context.Context, id int64) (err error)
	// Update the artifact. Only the properties specified by "props" will be updated if it is set
	Update(ctx context.Context, appTemplate *AppTemplate, props ...string) (err error)
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
	dao dao.DAO
}

func (m *manager) Count(ctx context.Context, query *q.Query) (total int64, err error) {
	return
}

// List apptemplate according to the query. The artifacts that referenced by others and
// without tags are not returned
func (m *manager) List(ctx context.Context, query *q.Query) (appTemplates []*AppTemplate, err error) {
	return
}
func (m *manager) Add(ctx context.Context, appTemplate *AppTemplate) (id int64, globalID string, err error) {
	id, globalID, err = m.dao.Create(ctx, appTemplate.To())
	if err != nil {
		return
	}
	return
}

// Get the apptemplate specified by the ID
func (m *manager) Get(ctx context.Context, id int64) (appTemplate *AppTemplate, err error) {
	return
}
func (m *manager) GetByGlobalID(ctx context.Context, globalid string) (appTemplate *AppTemplate, err error) {
	return
}

// Create the apptemplate. If the artifact is an index, make sure all the artifacts it references
// already exist
func (m *manager) Create(ctx context.Context, appTemplate *AppTemplate) (id int64, err error) {
	id, _, err = m.dao.Create(ctx, appTemplate.To())
	if err != nil {
		return 0, err
	}
	return
}

// Delete just deletes the artifact record. The underlying data of registry will be
// removed during garbage collection
func (m *manager) Delete(ctx context.Context, id int64) (err error) {
	return
}
func (m *manager) DeleteLogic(ctx context.Context, id int64) (err error) {

	return
}

// Update the artifact. Only the properties specified by "props" will be updated if it is set
func (m *manager) Update(ctx context.Context, appTemplate *AppTemplate, props ...string) (err error) {
	return
}
