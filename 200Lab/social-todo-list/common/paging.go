package common

type Paging struct {
	Page  int   `json:"page,omitempty" form:"page,omitempty"`
	Limit int   `json:"limit,omitempty" form:"limit,omitempty"`
	Total int64 `json:"total,omitempty" form:"-"`
}

func (p *Paging) Process() {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 {
		p.Limit = 1
	}

	if p.Limit >= 100 {
		p.Limit = 100
	}
}
