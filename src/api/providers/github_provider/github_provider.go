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


var headerAuthorization string

const (
	headerAuthorizationFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func init() {
	/*err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}*/

	// headerAuthorization = os.Getenv("TOKEN_GITHUB")
	headerAuthorization = "Authorization"
}


func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (
	*github.CreateRepoResponse,
	*github.GithubErrorResponse,
) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader("abc123"))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	fmt.Println("Github response", response)
	fmt.Println("Github err", err)

	if err != nil {
		log.Printf("error when trying to create a new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	// Validate body response
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "invalid response body",
		}
	}
	defer response.Body.Close()

	// Validate github error via status code
	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshal create repo successful response: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil
}