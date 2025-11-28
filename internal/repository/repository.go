package repository

type Author struct {
	ID   int
	Name string
}

type Repository interface {
	GetAuthorsByIDs(ids []int) ([]Author, error)
}
