package github_provider

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader (t *testing.T) {
	header := getAuthorizationHeader("abc123")

	assert.EqualValues(t, "token abc123", header)
}

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	defer fmt.Println("5")

	fmt.Println("function's body")
}