package hardcodedrepository

import (
	"author-service/internal/repository"
)

type HardcodedRepository struct {
	authors map[int]repository.Author
}

// NewHardcodedRepository creates a new hardcoded repository instance with data from quotes.csv
func NewHardcodedRepository() *HardcodedRepository {
	authors := map[int]repository.Author{
		1:  {ID: 1, Name: "Mark Twain"},
		2:  {ID: 2, Name: "Henry Ford"},
		3:  {ID: 3, Name: "Mark Twain"},
		4:  {ID: 4, Name: "Kurt Vonnegut"},
		5:  {ID: 5, Name: "Robert Frost"},
		6:  {ID: 6, Name: "Andrew Carnegie"},
		7:  {ID: 7, Name: "C. S. Lewis"},
		8:  {ID: 8, Name: "Confucius"},
		9:  {ID: 9, Name: "Eleanor Roosevelt"},
		10: {ID: 10, Name: "Samuel Ullman"},
		11: {ID: 11, Name: "Agatha Christie"},
		12: {ID: 12, Name: "Ralph Waldo Emerson"},
		13: {ID: 13, Name: "Aristotle"},
		14: {ID: 14, Name: "Bill Cosby"},
		15: {ID: 15, Name: "Francis Bacon"},
		16: {ID: 16, Name: "Henry David Thoreau"},
		17: {ID: 17, Name: "George Bernard Shaw"},
		18: {ID: 18, Name: "Victor Hugo"},
		19: {ID: 19, Name: "George Burns"},
		20: {ID: 20, Name: "Albert Camus"},
	}

	return &HardcodedRepository{
		authors: authors,
	}
}

// GetAuthorsByIDs returns authors matching the given IDs
func (r *HardcodedRepository) GetAuthorsByIDs(ids []int) ([]repository.Author, error) {
	var result []repository.Author

	for _, id := range ids {
		if author, ok := r.authors[id]; ok {
			result = append(result, author)
		}
	}

	return result, nil
}
