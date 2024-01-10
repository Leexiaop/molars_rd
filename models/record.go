package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Record struct {
	Model

	Price     int    `json:"price"`
	ProductId int    `json:"product_id"`
	Count     int    `json:"count"`
	Url       string `json:"url"`
}

func GetRecords(pageNum int, pageSize int, maps interface{}) (records []Record) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&records)
	return
}

func GetRecordsTotal(maps interface{}) (count int) {
	db.Model(&Record{}).Where(maps).Count(&count)
	return
}

func AddRecords (price int, count int, productId int, url string) bool {
	db.Create(&Record{
		Price: price,
		Count: count,
		ProductId: productId,
		Url: url,
	})
	return true
}

func ExistRecordId (id int) bool {
	var record Record
	db.Select("id").Where("id = ?", id).First(&record)
	return record.ID > 0
}

func EditRecords (id int, data interface{}) bool {
	db.Model(&Record{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteRecords (id int) bool {
	db.Where("id = ?", id).Delete(&Record{})
	return true
}

func (product *Record) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	scope.SetColumn("CreatedBy", "13691388204")
	return nil
}

func (product *Record) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifieldOn", time.Now().Unix())
	scope.SetColumn("ModifieldBy", "13691388204")
	return nil
}