package apptemplate

import (
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor/src/lib/log"
	"github.com/goharbor/harbor/src/pkg/apptemplate/dao"
)

// AppTemplate is an abstract object managed by Harbor

type AppTemplate struct {
	ID          int64                  `json:"id"`
	GlobalID    string                 `json:"global_id"`
	AppName     string                 `json:"app_name"`
	Icon        string                 `json:"icon"`
	Desc        string                 `json:"desc"`
	ProjectID   int64                  `json:"project_id"`
	ManifestLoc string                 `json:"manifest_loc"`
	UpdateTime  string                 `json:"update_time"`
	CreateTime  string                 `json:"create_time"`
	Status      int                    `json:"status"`
	ExtraAttrs  map[string]interface{} `json:"extra_attrs"` //json
	Annotations map[string]string      `json:"annotations"` //json
}

func (at *AppTemplate) String() string {
	return fmt.Sprintf("%s@%s", at.AppName, at.GlobalID)
}

func (at *AppTemplate) From(atdao *dao.AppTemplate) {
	at.ID = atdao.ID
	at.GlobalID = atdao.GlobalID
	at.AppName = atdao.AppName
	at.Icon = atdao.ICON
	at.Desc = atdao.Desc
	at.ManifestLoc = atdao.ManifestLoc
	at.ProjectID = atdao.ProjectID
	at.UpdateTime = atdao.UpdateTime
	at.CreateTime = atdao.CreateTime
	at.Status = atdao.Status
	at.ExtraAttrs = map[string]interface{}{}
	at.Annotations = map[string]string{}
	if len(atdao.ExtraAttrs) > 0 {
		if err := json.Unmarshal([]byte(atdao.ExtraAttrs), &at.ExtraAttrs); err != nil {
			log.Errorf("failed to unmarshal extra attr%d,%s", atdao.ID, atdao.GlobalID)
		}
	}
	if len(atdao.Annotations) > 0 {
		if err := json.Unmarshal([]byte(atdao.Annotations), &at.Annotations); err != nil {
			log.Errorf("failed to unmarshal the annotations of artifact %d: %v", atdao.ID, err)
		}
	}
	return
}

func (at *AppTemplate) To() *dao.AppTemplate {
	atdao := &dao.AppTemplate{
		ID:          at.ID,
		GlobalID:    at.GlobalID,
		ManifestLoc: at.ManifestLoc,
		ProjectID:   at.ProjectID,
		ICON:        at.Icon,
		UpdateTime:  at.UpdateTime,
		CreateTime:  at.CreateTime,
	}
	if len(at.ExtraAttrs) > 0 {
		attrs, err := json.Marshal(at.ExtraAttrs)
		if err != nil {
			log.Errorf("failed to unmarshal extra attr%d,%s", atdao.ID, atdao.GlobalID)
		}
		atdao.ExtraAttrs = string(attrs)
	}
	if len(at.Annotations) > 0 {
		annotations, err := json.Marshal(at.Annotations)
		if err != nil {
			log.Errorf("failed to unmarshal the annotations of artifact %d: %v", atdao.ID, err)
		}
		atdao.Annotations = string(annotations)
	}
	return atdao
}
