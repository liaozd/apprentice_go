package data

import "greenlight.alexedwards.net/internal/validator"

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func ValidateFilters(v *validator.Validator, f Filters) {
	// curl "localhost:4000/v1/movies?page=-1&page_size=-1&sort=foo"
	//{
	//	"error": {
	//	"page": "must be greater than zero",
	//	"page_size": "must be greater than zero",
	//	"sort": "invalid sort value"
	//	}
	//}
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 1-million")

	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")

	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
