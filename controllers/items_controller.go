package controllers

import (
	"fmt"
	"net/http"

	"github.com/shchaslyvyi/bookstore_items_api/domain/items"
	"github.com/shchaslyvyi/bookstore_items_api/services"
	"github.com/shchaslyvyi/bookstore_oauth-go/oauth"
)

//
var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

// Create func creates the item
func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error json to the caller.
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemsService.Create(item)
	fmt.Println(err)
	//if err != nil {
	// TODO: Return error json to the caller.
	//}

	fmt.Println(result)
	// TODO: Return created item with HTTP status 201 - Created.
}

// Get func gets the item
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}
