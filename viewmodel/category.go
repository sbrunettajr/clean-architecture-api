package viewmodel

type CategoryRequestVM struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type CategoryResponseVM struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}
