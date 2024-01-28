package record_service

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/export"
	"github.com/Leexiaop/molars_rd/pkg/file"
	"github.com/Leexiaop/molars_rd/pkg/gredis"
	"github.com/Leexiaop/molars_rd/pkg/logging"
	"github.com/Leexiaop/molars_rd/service/cache_service"
	"github.com/tealeg/xlsx"
)

type Record struct {
	ID         int
	ProductId  int
	Price      int
	Count      int
	CreatedBy  string
	Url        string
	ModifiedBy string
	PageNum    int
	PageSize   int
	Name string
}

func (r *Record) Counts() (int, error) {
	return models.GetRecordsTotal(r.getMaps())
}
func (r*Record) GetAll () ([]models.Record, error) {
	var (
		records, cacheRecords []models.Record
	)

	cache := cache_service.Record{
		ProductId: r.ProductId,
		PageSize: r.PageSize,
		PageNum: r.PageNum,
	}
	key := cache.GetRecordsKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheRecords)
			return cacheRecords, nil
		}
	}
	records, err := models.GetRecords(r.PageNum, r.PageSize, r.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, records, 3600)
	return records, nil
}
func (r *Record) getMaps() map[string]interface{}  {
	maps := make(map[string]interface{})
	if r.ProductId != 0 {
		maps["product_id"] = r.ProductId
	}
	return maps
}

func (r *Record) ExistById() (bool, error) {
	return models.ExistRecordId(r.ID)
}

func (r *Record) Delete() error {
	return models.DeleteRecord(r.ID)
}
func (r *Record) Add() error {
	return models.AddRecords(r.Price, r.Count, r.ProductId, r.Url)
}
func (r *Record) Edit() error {
	return models.EditRecords(r.ID, map[string]interface{}{
		"price": r.Price,
		"count": r.Count,
		"product_id": r.ProductId,
		"url": r.Url,
	})
}
func (r *Record) Export() (string, error) {
	records, err := r.GetAll()
	if err != nil {
		return "", err
	}

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("记录信息")

	if err != nil {
		return "", err
	}
	titles := []string{"ID", "价格", "数量", "产品ID","创建时间","创建人","修改人","修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell

	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range records {
		values := []string{
			strconv.Itoa(v.ID),
			strconv.Itoa(v.Price),
			strconv.Itoa(v.Count),
			strconv.Itoa(v.ProductId),
			v.CreatedBy,
			v.ModifieldBy,
			strconv.Itoa(v.CreatedOn),
			strconv.Itoa(v.ModifieldOn),
		}
		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time.Now().Unix()))
	filename := "record_" + time + ".xlsx"

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}
	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}