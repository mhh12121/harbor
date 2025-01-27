// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	"github.com/goharbor/harbor/src/lib/config"
	"github.com/goharbor/harbor/src/pkg/artifact"
	cachedArtifact "github.com/goharbor/harbor/src/pkg/cached/artifact/redis"
	cachedProject "github.com/goharbor/harbor/src/pkg/cached/project/redis"
	"github.com/goharbor/harbor/src/pkg/project"
)

// Define global resource manager.
var (
	// ArtifactMgr is the manager for artifact.
	ArtifactMgr artifact.Manager
	// ProjectMgr is the manager for project.
	ProjectMgr project.Manager
)

// init initialize mananger for resources
func init() {
	cacheEnabled := config.CacheEnabled()
	initArtifactMgr(cacheEnabled)
	initProjectMgr(cacheEnabled)

}

func initArtifactMgr(cacheEnabled bool) {
	artMgr := artifact.NewManager()
	// check cache enable
	if cacheEnabled {
		ArtifactMgr = cachedArtifact.NewManager(artMgr)
	} else {
		ArtifactMgr = artMgr
	}
}

func initProjectMgr(cacheEnabled bool) {
	projectMgr := project.New()
	// check cache enable
	if cacheEnabled {
		ProjectMgr = cachedProject.NewManager(projectMgr)
	} else {
		ProjectMgr = projectMgr
	}
}
