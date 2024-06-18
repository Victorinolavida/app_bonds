package data

import "boundsApp.victorinolavida/internal/validator"

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records"`
}

func (p *Pagination) offset() int {
	return (p.CurrentPage - 1) * p.PageSize
}

func (p *Pagination) limit() int {
	return p.PageSize
}

func ValidatePagination(v *validator.Validator, p *Pagination) {
	v.Check(p.CurrentPage > 0, "current_page", "must be greater than zero")
	v.Check(p.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(p.PageSize <= 100, "page_size", "must be less than or equal to 100")

}

func getPagination(totalRecords, page, pageSize int) Pagination {
	return Pagination{
		CurrentPage:  page,
		PageSize:     pageSize,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}
