package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "a golang introduction",
		Description: "a golang introduction repository",
		HomePage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, `{"name":"a golang introduction","description":"a golang introduction repository","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`,
		// Ensure input string matches bytes generates
		string(bytes),
	)

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}
