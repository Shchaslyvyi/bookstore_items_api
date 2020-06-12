package services

import (
	"net/http"

	"github.com/shchaslyvyi/bookstore_items_api/domain/items"
	"github.com/shchaslyvyi/bookstore_utils-go/rest_errors"
)

// ItemsService ot type ItemsServiceInterface
var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.RestErr{"to be implemented!", http.StatusNotImplemented, "not_implemented", nil}
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.RestErr{"to be implemented!", http.StatusNotImplemented, "not_implemented", nil}
}
