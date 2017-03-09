package client

import (
	"github.com/michaellihs/golab/model"
)

type ProjectsService struct {
	Client *GitlabClient
}

type ProjectParams struct {
	Name                                             string `json:"name"`
	path                                             string
	namespace_id                                     int
	default_branch                                   string
	description                                      string
	issues_enabled                                   bool
	merge_requests_enabled                           bool
	builds_enabled                                   bool
	wiki_enabled                                     bool
	snippets_enabled                                 bool
	container_registry_enabled                       bool
	shared_runners_enabled                           bool
	visibility                                       string
	import_url                                       string
	public_builds                                    bool
	only_allow_merge_if_pipeline_succeeds            bool
	only_allow_merge_if_all_discussions_are_resolved bool
	lfs_enabled                                      bool
	request_access_enabled                           bool
	repository_storage                               string
	approvals_before_merge                           int
}

func (service *ProjectsService) List() *[]model.Project {
	projects := new([]model.Project)
	req, _ := service.Client.NewGetRequest("/api/v3/projects")
	service.Client.Do(req, projects)
	return projects
}

func (service *ProjectsService) Create(projectParams *ProjectParams) *model.Project {
	req, _ := service.Client.NewPostRequest("/api/v3/projects", projectParams)
	project := new(model.Project)
	service.Client.Do(req, project)
	return project
}