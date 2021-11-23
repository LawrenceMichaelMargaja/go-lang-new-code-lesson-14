package github_provider_test

import (
	"fmt"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/domain/github"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/providers/github_provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := github_provider.GetAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestDefer(t *testing.T)  {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("function's body")
}

func TestCreateRepo(t *testing.T) {
	github_provider.CreateRepo("", github.CreateRepoRequest{})
}