package services

import (
	"github.com/dembygenesis/go-rest-industry-standard/mvc/domain"
	"github.com/dembygenesis/go-rest-industry-standard/mvc/utils"
	"net/http"
)

type itemsService struct {

}

func (s *itemsService) GetItem (itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "",
	}
}