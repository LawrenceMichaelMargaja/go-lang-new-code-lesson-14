package repositories

import (
	"encoding/json"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/repositories"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/services"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/utils/errors"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


var (
	funcCreateRepo func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
)

type repoServiceMock struct {

}

func (s *repoServiceMock) CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo("", request)
}

func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	// Override the function's public variable during testing
	services.RepositoryService = &repoServiceMock{}

	// Override mock function
	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			Id:    321,
			Owner: "golang",
			Name:  "mocked service",
		}, nil
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, response)

	// Execute gin controller function
	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.EqualValues(t, 321, result.Id)
	assert.EqualValues(t, "golang", result.Owner)
	assert.EqualValues(t, "mocked service", result.Name)
}

func TestCreateRepoErrorMockingTheEntireService(t *testing.T) {
	// Override the function's public variable during testing
	services.RepositoryService = &repoServiceMock{}

	// Override mock function
	funcCreateRepo = func(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	c := test_utils.GetMockedContext(request, response)

	// Execute gin controller function
	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid repository name", apiErr.Message())
}