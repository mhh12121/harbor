package apptemplate

import (
	"context"
	"github.com/goharbor/harbor/src/lib/q"
	"github.com/goharbor/harbor/src/pkg/apptemplate"
)

var (
	Ctl = NewController()
)

type Controller interface {
	// Ensure the artifact specified by the digest exists under the repository,
	// creates it if it doesn't exist. If tags are provided, ensure they exist
	// and are attached to the artifact. If the tags don't exist, create them first.
	// The "created" will be set as true when the artifact is created
	Ensure(ctx context.Context, repository string, globalID string) (created bool, id int64, err error)
	// Count returns the total count of artifacts according to the query.
	// The artifacts that referenced by others and without tags are not counted
	Count(ctx context.Context, query *q.Query) (total int64, err error)
	// List artifacts according to the query, specify the properties returned with option
	// The artifacts that referenced by others and without tags are not returned
	List(ctx context.Context, query *q.Query, option *Option) (appTemplates []*AppTemplate, err error)
	// Get the artifact specified by ID, specify the properties returned with option
	Get(ctx context.Context, globalID int64, option *Option) (appTemplate *AppTemplate, err error)
	Add(ctx context.Context, appTemplate *AppTemplate) (err error)
	// Delete the artifact specified by apptemplate ID
	Delete(ctx context.Context, globalID int64) (err error)
	// Copy the artifact specified by "srcRepo" and "reference" into the repository specified by "dstRepo"
	Copy(ctx context.Context, srcRepo, reference, dstRepo string) (id int64, err error)
	// // UpdatePullTime updates the pull time for the artifact. If the tagID is provides, update the pull
	// // time of the tag as well
	// UpdatePullTime(ctx context.Context, artifactID int64, tagID int64, time time.Time) (err error)
	// // GetAddition returns the addition of the artifact.
	// // The addition is different according to the artifact type:
	// // build history for image; values.yaml, readme and dependencies for chart, etc
	// GetAddition(ctx context.Context, artifactID int64, additionType string) (addition *processor.Addition, err error)
	// // AddLabel to the specified artifact
	// AddLabel(ctx context.Context, artifactID int64, labelID int64) (err error)
	// // RemoveLabel from the specified artifact
	// RemoveLabel(ctx context.Context, artifactID int64, labelID int64) (err error)
	// Walk walks the artifact tree rooted at root, calling walkFn for each artifact in the tree, including root.
	Walk(ctx context.Context, root *AppTemplate, walkFn func(*AppTemplate) error, option *Option) error
}

func NewController() Controller {
	return &controller{
		apptMgr: apptemplate.Mgr,
	}
}

type controller struct {
	apptMgr apptemplate.Manager
}

func (c *controller) Ensure(ctx context.Context, repository string, globalID string) (created bool, id int64, err error) {
	return
}

// Count returns the total count of artifacts according to the query.
// The artifacts that referenced by others and without tags are not counted
func (c *controller) Count(ctx context.Context, query *q.Query) (total int64, err error) {
	return
}

// List artifacts according to the query, specify the properties returned with option
// The artifacts that referenced by others and without tags are not returned
func (c *controller) List(ctx context.Context, query *q.Query, option *Option) (appTemplates []*AppTemplate, err error) {
	return
} // Get the artifact specified by ID, specify the properties returned with option
func (c *controller) Get(ctx context.Context, globalID int64, option *Option) (appTemplate *AppTemplate, err error) {
	return
}
func (c *controller) Add(ctx context.Context, appTemplate *AppTemplate) (err error) {
	// c.apptMgr
	return
}

// Delete the artifact specified by apptemplate ID
func (c *controller) Delete(ctx context.Context, globalID int64) (err error) {
	return
}

// Copy the artifact specified by "srcRepo" and "reference" into the repository specified by "dstRepo"
func (c *controller) Copy(ctx context.Context, srcRepo, reference, dstRepo string) (id int64, err error) {
	return
}

// // UpdatePullTime updates the pull time for the artifact. If the tagID is provides, update the pull
// // time of the tag as well
// UpdatePullTime(ctx context.Context, artifactID int64, tagID int64, time time.Time) (err error)
// // GetAddition returns the addition of the artifact.
// // The addition is different according to the artifact type:
// // build history for image; values.yaml, readme and dependencies for chart, etc
// GetAddition(ctx context.Context, artifactID int64, additionType string) (addition *processor.Addition, err error)
// // AddLabel to the specified artifact
// AddLabel(ctx context.Context, artifactID int64, labelID int64) (err error)
// // RemoveLabel from the specified artifact
// RemoveLabel(ctx context.Context, artifactID int64, labelID int64) (err error)
// Walk walks the artifact tree rooted at root, calling walkFn for each artifact in the tree, including root.
func (c *controller) Walk(ctx context.Context, root *AppTemplate, walkFn func(*AppTemplate) error, option *Option) (err error) {
	return
}
