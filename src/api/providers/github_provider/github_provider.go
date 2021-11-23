package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/clients/restclient"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func GetAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// access token = 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

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, GetAuthorizationHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)

	fmt.Println(response)
	fmt.Println(err)

	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create a new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "Invalid Response Body.",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "Invalid JSON body.",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "error when trying to unmarshal github create repo response",
		}
	}

 	return &result, nil
}