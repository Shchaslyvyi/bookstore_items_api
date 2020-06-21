package items

import (
	"errors"

	"github.com/shchaslyvyi/bookstore_items_api/clients/elasticsearch"
	"github.com/shchaslyvyi/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

// Save func
func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("DB error"))
	}
	i.ID = result.Id
	return nil
}
