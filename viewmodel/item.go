package viewmodel

type ItemRequest struct {
	Active       bool    `json:"active"`
	CategoryUUID string  `json:"category_uuid"`
	Description  string  `json:"description"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
}

type ItemResponse struct {
	Active      bool               `json:"active"`
	Category    CategoryResponseVM `json:"category"`
	Description string             `json:"description"`
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
}
