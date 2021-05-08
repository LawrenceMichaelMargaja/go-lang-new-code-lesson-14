package services

import (
	"github.com/dembygenesis/go-rest-industry-standard/src/api/config"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/github"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/repositories"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/log/option_b"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/providers/github_provider"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/utils/errors"
	"net/http"
	"sync"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(clientId string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
	/*handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse)
	createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult)*/
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	// Validate name
	if err := input.Validate(); err != nil {
		return nil, err
	}

	// Create request object
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	option_b.Info("about to send request to external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "pending"))

	// Execute api call
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		option_b.Info("response obtained from external api",
			option_b.Field("client_id", clientId),
			option_b.Field("status", "error"))
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	// Return result
	option_b.Info("response obtained from external api",
		option_b.Field("client_id", clientId),
		option_b.Field("status", "success"))

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	var wg sync.WaitGroup

	go s.handleRepoResults(&wg, input, output)
	for _, current := range request {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <- output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}

	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *reposService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}


	output <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo("SAMPLE_CLIENT_ID", input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}
