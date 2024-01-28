package cache_service

import (
	"strconv"
	"strings"

	"github.com/Leexiaop/molars_rd/pkg/e"
)

type Product struct {
	ID       int
	PageNum  int
	PageSize int
}

func (p *Product) GetProductsKey() string {
	keys := []string{
		e.CACHE_PRODUCT,
		"LIST",
	}

	if p.PageNum > 0 {
		keys = append(keys, strconv.Itoa(p.PageNum))
	}
	if p.PageSize > 0 {
		keys = append(keys, strconv.Itoa(p.PageSize))
	}

	return strings.Join(keys, "_")
}