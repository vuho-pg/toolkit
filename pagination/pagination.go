package pagination

const (
	DefaultLimit = 20
	DefaultPage  = 1
)

type Pagination struct {
	ok          bool
	noLimit     bool
	noOrder     bool
	Page        int     `json:"page"`
	TotalPage   int64   `json:"total_page"`
	TotalRecord int64   `json:"total_record"`
	Limit       int     `json:"limit"`
	Sort        string  `json:"order"`
	Orders      []Order `json:"-"`
}

func (p *Pagination) NoLimit(noLimit bool) *Pagination {
	p.noLimit = noLimit
	return p
}
func (p *Pagination) NoOrder(noOrder bool) *Pagination {
	p.noOrder = noOrder
	return p
}

func (p *Pagination) correct() {
	if p.Page < 1 {
		p.Page = DefaultPage
	}
	if p.Limit == 0 {
		p.Limit = DefaultLimit
	}
}

func (p *Pagination) init() {
	p.Orders = splitOrder(p.Sort)
	p.ok = true
}

func (p *Pagination) SetTotalRecord(total int64) {
	p.TotalRecord = total
	p.TotalPage = total / int64(p.Limit)
	if total%int64(p.Limit) != 0 {
		p.TotalPage++
	}
}
