package model_test

import (
	. "github.com/michaellihs/golab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"encoding/json"
)

var _ = Describe("Group", func() {

	It("can be unmarshalled from sample JSON", func() {
		group := new(Group)
		// sample taken from Gitlab API doc: https://docs.gitlab.com/ee/api/groups.html
		json.Unmarshal([]byte(`{
			  "id": 4,
			  "name": "Twitter",
			  "path": "twitter",
			  "description": "Aliquid qui quis dignissimos distinctio ut commodi voluptas est.",
			  "visibility": "public",
			  "avatar_url": null,
			  "web_url": "https://gitlab.example.com/groups/twitter",
			  "request_access_enabled": false,
			  "full_name": "Twitter",
			  "full_path": "twitter",
			  "parent_id": null,
			  "projects": [
			    {
			      "id": 7,
			      "description": "Voluptas veniam qui et beatae voluptas doloremque explicabo facilis.",
			      "default_branch": "master",
			      "tag_list": [],
			      "archived": false,
			      "visibility": "public",
			      "ssh_url_to_repo": "git@gitlab.example.com:twitter/typeahead-js.git",
			      "http_url_to_repo": "https://gitlab.example.com/twitter/typeahead-js.git",
			      "web_url": "https://gitlab.example.com/twitter/typeahead-js",
			      "name": "Typeahead.Js",
			      "name_with_namespace": "Twitter / Typeahead.Js",
			      "path": "typeahead-js",
			      "path_with_namespace": "twitter/typeahead-js",
			      "issues_enabled": true,
			      "merge_requests_enabled": true,
			      "wiki_enabled": true,
			      "jobs_enabled": true,
			      "snippets_enabled": false,
			      "container_registry_enabled": true,
			      "created_at": "2016-06-17T07:47:25.578Z",
			      "last_activity_at": "2016-06-17T07:47:25.881Z",
			      "shared_runners_enabled": true,
			      "creator_id": 1,
			      "namespace": {
				"id": 4,
				"name": "Twitter",
				"path": "twitter",
				"kind": "group"
			      },
			      "avatar_url": null,
			      "star_count": 0,
			      "forks_count": 0,
			      "open_issues_count": 3,
			      "public_jobs": true,
			      "shared_with_groups": [],
			      "request_access_enabled": false
			    },
			    {
			      "id": 6,
			      "description": "Aspernatur omnis repudiandae qui voluptatibus eaque.",
			      "default_branch": "master",
			      "tag_list": [],
			      "archived": false,
			      "visibility": "internal",
			      "ssh_url_to_repo": "git@gitlab.example.com:twitter/flight.git",
			      "http_url_to_repo": "https://gitlab.example.com/twitter/flight.git",
			      "web_url": "https://gitlab.example.com/twitter/flight",
			      "name": "Flight",
			      "name_with_namespace": "Twitter / Flight",
			      "path": "flight",
			      "path_with_namespace": "twitter/flight",
			      "issues_enabled": true,
			      "merge_requests_enabled": true,
			      "wiki_enabled": true,
			      "jobs_enabled": true,
			      "snippets_enabled": false,
			      "container_registry_enabled": true,
			      "created_at": "2016-06-17T07:47:24.661Z",
			      "last_activity_at": "2016-06-17T07:47:24.838Z",
			      "shared_runners_enabled": true,
			      "creator_id": 1,
			      "namespace": {
				"id": 4,
				"name": "Twitter",
				"path": "twitter",
				"kind": "group"
			      },
			      "avatar_url": null,
			      "star_count": 0,
			      "forks_count": 0,
			      "open_issues_count": 8,
			      "public_jobs": true,
			      "shared_with_groups": [],
			      "request_access_enabled": false
			    }
			  ],
			  "shared_projects": [
			    {
			      "id": 8,
			      "description": "Velit eveniet provident fugiat saepe eligendi autem.",
			      "default_branch": "master",
			      "tag_list": [],
			      "archived": false,
			      "visibility": "private",
			      "ssh_url_to_repo": "git@gitlab.example.com:h5bp/html5-boilerplate.git",
			      "http_url_to_repo": "https://gitlab.example.com/h5bp/html5-boilerplate.git",
			      "web_url": "https://gitlab.example.com/h5bp/html5-boilerplate",
			      "name": "Html5 Boilerplate",
			      "name_with_namespace": "H5bp / Html5 Boilerplate",
			      "path": "html5-boilerplate",
			      "path_with_namespace": "h5bp/html5-boilerplate",
			      "issues_enabled": true,
			      "merge_requests_enabled": true,
			      "wiki_enabled": true,
			      "jobs_enabled": true,
			      "snippets_enabled": false,
			      "container_registry_enabled": true,
			      "created_at": "2016-06-17T07:47:27.089Z",
			      "last_activity_at": "2016-06-17T07:47:27.310Z",
			      "shared_runners_enabled": true,
			      "creator_id": 1,
			      "namespace": {
				"id": 5,
				"name": "H5bp",
				"path": "h5bp",
				"kind": "group"
			      },
			      "avatar_url": null,
			      "star_count": 0,
			      "forks_count": 0,
			      "open_issues_count": 4,
			      "public_jobs": true,
			      "shared_with_groups": [
				{
				  "group_id": 4,
				  "group_name": "Twitter",
				  "group_access_level": 30
				},
				{
				  "group_id": 3,
				  "group_name": "Gitlab Org",
				  "group_access_level": 10
				}
			      ]
			    }
			  ]
			}`), group)
		Expect(group.Name).To(Equal("Twitter"))
		Expect(group.Path).To(Equal("twitter"))
		Expect(group.Projects[0].Name).To(Equal("Typeahead.Js"))
	})

	It("can be unmarshalled from JSON without projects", func() {
		group := new(Group)
		json.Unmarshal([]byte(`{
			"id": 4,
			"name": "Twitter",
			"path": "twitter",
			"description": "Aliquid qui quis dignissimos distinctio ut commodi voluptas est.",
			"visibility": "public",
			"avatar_url": null,
			"web_url": "https://gitlab.example.com/groups/twitter",
			"request_access_enabled": false,
			"full_name": "Twitter",
			"full_path": "twitter",
			"parent_id": null,
			"projects": null,
			"shared_projects": null
		}`), group)

		Expect(group.Name).To(Equal("Twitter"))
	})
})