package items

// Item struct to define the Item object
type Item struct {
	ID                string      `json:"ID"`
	Seller            int64       `json:"seller"`
	Title             string      `json:"title"`
	Decription        Description `json:"description"`
	Pictures          []Picture   `json:"picture"`
	Video             string      `json:"video"`
	Price             float32     `json:"price"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Status            string      `json:"status"`
}

// Description nested scruct to define the Item description
type Description struct {
	PlainText string `json:"plain_text"`
	HTML      string `json:"html"`
}

// Picture nested scruct to define the Picture object of the Item
type Picture struct {
	ID  int64  `json:"ID"`
	URL string `json:"URL"`
}
