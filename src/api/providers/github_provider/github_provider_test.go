package github_provider

/**
Tests our native http client.
I'm going to assume we are overriding our client
*/

import (
	"errors"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/clients/restclient"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

// This is the entry point of tests
// This function calls each function prefixed with "Test" and runs them
func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")

	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	// Modifies the state of "enabledMock" in the rest client lib to "true"
	// Called at main entry instead
	// restclient.StartMockups()

	// Start fresh
	restclient.FlushMockups()

	// Add a mock struct to the response
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid rest client response"),
	})

	// Execute, but take note "enabledMock" is true - hence this will
	// not perform actual rest calls
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid rest client response", err.Message)
}

func TestCreateRepoErrorInvalidResponseBody(t *testing.T) {
	// Modifies the state of "enabledMock" in the rest client lib to "true"
	// Called at main entry instead
	// restclient.StartMockups()

	// Start fresh
	restclient.FlushMockups()

	// Generates an invalid body
	invalidCloser, _ := os.Open("-asf3")

	// Add a mock struct to the response
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			// Key knowledge: Generates a response body error
			Body: invalidCloser,
		},
	})

	// Execute, but take note "enabledMock" is true - hence this will
	// not perform actual rest calls
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(t *testing.T) {
	// Modifies the state of "enabledMock" in the rest client lib to "true"
	// Called at main entry instead
	// restclient.StartMockups()

	// Start fresh
	restclient.FlushMockups()

	// Add a mock struct to the response
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			// Key knowledge: Generates a json response error
			Body: ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	// Execute, but take note "enabledMock" is true - hence this will
	// not perform actual rest calls
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response", err.Message)
}

func TestCreateRepoUnauthorized(t *testing.T) {
	// Modifies the state of "enabledMock" in the rest client lib to "true"
	// Called at main entry instead
	// restclient.StartMockups()

	// Start fresh
	restclient.FlushMockups()

	// Add a mock struct to the response
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			// Key knowledge: Generates a json response error
			Body: ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://docs.github.com/rest/reference/repos#create-a-repository-for-the-authenticated-user"}`)),
		},
	})

	// Execute, but take note "enabledMock" is true - hence this will
	// not perform actual rest calls
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

func TestCreateRepoInvalidSuccessResponse(t *testing.T) {
	// Modifies the state of "enabledMock" in the rest client lib to "true"
	// Called at main entry instead
	// restclient.StartMockups()

	// Start fresh
	restclient.FlushMockups()

	// Add a mock struct to the response
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})

	// Execute, but take note "enabledMock" is true - hence this will
	// not perform actual rest calls
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal github create repo response", err.Message)
}

func TestCreateRepoNoError(t *testing.T) {
	// Modifies the state of "enabledMock" in the rest client lib to "true"
	// Called at main entry instead
	// restclient.StartMockups()

	// Start fresh
	restclient.FlushMockups()

	// Add a mock struct to the response
	restclient.AddMockup(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": ""}`)),
		},
	})

	// Execute, but take note "enabledMock" is true - hence this will
	// not perform actual rest calls
	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "", response.Name)
	assert.EqualValues(t, "", response.FullName)
}
