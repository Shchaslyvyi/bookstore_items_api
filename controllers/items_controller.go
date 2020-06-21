package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shchaslyvyi/bookstore_items_api/domain/items"
	"github.com/shchaslyvyi/bookstore_items_api/services"
	"github.com/shchaslyvyi/bookstore_items_api/utils/http_utils"
	"github.com/shchaslyvyi/bookstore_oauth-go/oauth"
	"github.com/shchaslyvyi/bookstore_utils-go/rest_errors"
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
		//http_utils.RespondError(w, *err)
		return
	}

	sellerID := oauth.GetCallerID(r)

	if sellerID == 0 {
		respErr := rest_errors.NewNotAuthorizedError("invalid request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, *restErr)
		return
	}

	itemRequest.Seller = sellerID

	result, createErr := services.ItemsService.Create(itemRequest)
	fmt.Println(createErr)
	if err != nil {
		http_utils.RespondError(w, *createErr)
		return
	}
	http_utils.RespondJSON(w, http.StatusCreated, result)
}

// Get func gets the item
func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}
