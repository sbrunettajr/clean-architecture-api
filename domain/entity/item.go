package entity

type Item struct {
	ID          int
	UUID        string
	Active      bool
	Category    Category
	CategoryID  int
	Description string
	Name        string
	Price       float64
}
