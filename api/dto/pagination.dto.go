package dto

type PaginationReq struct {
	Page    uint              `json:"page"`
	Perpage uint              `json:"perpage"`
	Filter  map[string]string `json:"filter"`
	Sort    map[string]string `json:"sort"`
}

func NewPaginationReq(in PaginationReq) PaginationReq {
	if in.Page <= 0 {
		in.Page = 1
	}

	if in.Perpage <= 0 {
		in.Perpage = 10
	}

	if in.Filter == nil {
		in.Filter = make(map[string]string)
	}

	if in.Sort == nil {
		in.Sort = make(map[string]string)
	}

	return PaginationReq{
		Page:    in.Page,
		Perpage: in.Perpage,
		Filter:  in.Filter,
		Sort:    in.Sort,
	}
}

func (p PaginationReq) Offset() uint {
	return (p.Page * p.Perpage) - p.Perpage
}

type PaginationRes[T any] struct {
	*PaginationReq
	Count uint `json:"count"`
	Items []T  `json:"items"`
}

func NewPaginationRes[T any](in PaginationRes[T]) PaginationRes[T] {
	if in.Items == nil {
		in.Items = make([]T, 0)
	}

	if in.Count < 0 {
		in.Count = 0
	}

	return PaginationRes[T]{
		PaginationReq: in.PaginationReq,
		Items:         in.Items,
		Count:         in.Count,
	}
}
