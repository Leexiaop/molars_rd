package cache_service

import (
	"strconv"
	"strings"

	"github.com/Leexiaop/molars_rd/pkg/e"
)

type Record struct {
	ID        int
	ProductId int
	PageSize  int
	PageNum   int
}

func (r *Record) GetRecordKey() string {
	return e.CACHE_RECORD + "_" + strconv.Itoa(r.ID)
}

func (r *Record) GetRecordsKey() string {
	keys := []string{
		e.CACHE_RECORD,
		"LIST",
	}
	if r.ProductId > 0 {
		keys = append(keys, strconv.Itoa(r.ProductId))
	}
	if r.PageNum > 0 {
		keys = append(keys, strconv.Itoa(r.PageNum))
	}
	if r.PageSize > 0 {
		keys = append(keys, strconv.Itoa(r.PageSize))
	}
	return strings.Join(keys, "_")
}