package pagination

const (
	DefaultLimit = 20
	DefaultPage  = 1
)

// Pagination : pagination utility.
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

// NoLimit : set to true if you want query has no limit.
func (p *Pagination) NoLimit(noLimit bool) *Pagination {
	p.noLimit = noLimit
	return p
}

// NoOrder : set to true if you don't want your result return in sorted order.
func (p *Pagination) NoOrder(noOrder bool) *Pagination {
	p.noOrder = noOrder
	return p
}

// Offset : return the offset value
func (p *Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
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
	p.correct()
	p.Orders = splitOrder(p.Sort)
	p.ok = true
}

func (p *Pagination) checkOk() {
	if p.ok {
		return
	}
	p.init()
}

// SetTotalRecord : use this will recalculate the TotalPage
func (p *Pagination) SetTotalRecord(total int64) {
	p.checkOk()
	p.TotalRecord = total
	p.TotalPage = total / int64(p.Limit)
	if total%int64(p.Limit) != 0 {
		p.TotalPage++
	}
}
