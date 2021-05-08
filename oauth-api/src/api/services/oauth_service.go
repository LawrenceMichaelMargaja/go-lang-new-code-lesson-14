package services

import (
	"github.com/dembygenesis/go-rest-industry-standard/oauth-api/src/api/controllers/domain/oauth"
	"github.com/dembygenesis/go-rest-industry-standard/src/api/utils/errors"
	"time"
)

type oauthService struct {}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	token := &oauth.AccessToken{
		AccessToken: "1234567890",
		UserId:      user.Id,
		Expires:     time.Now().UTC().Add(24 * time.Hour).Unix(),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}

	// TODO: Hard coded for now
	return token, nil
}

func (s* oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}
	return token, err
}
