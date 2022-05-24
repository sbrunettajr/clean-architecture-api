package entity

type Category struct {
	ID     int
	UUID   string
	Active bool
	Name   string
}

func (e Category) IsActive() bool {
	return e.Active
}
