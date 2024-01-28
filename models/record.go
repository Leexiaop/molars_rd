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

func GetRecords(pageNum int, pageSize int, maps interface{}) ([]Record, error) {
	var (
		records []Record
		err error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&records).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&records).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return records, nil
}

func GetRecordsTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Record{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func AddRecords (price int, count int, productId int, url string) error {
	record := Record{
		Price: price,
		Count: count,
		ProductId: productId,
		Url: url,
	}
	if err := db.Create(&record).Error; err != nil {
		return err
	}
	return nil
}

func ExistRecordId (id int) (bool, error) {
	var record Record
	err := db.Select("id").Where("id = ?", id).First(&record).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if record.ID > 0 {
		return true, nil
	}
	return false, nil
}

func EditRecords (id int, data interface{}) error {
	if err := db.Model(&Record{}).Where("id = ?", id).Updates(data).Error;err != nil {
		return err
	}
	return nil
}

func DeleteRecord (id int) error {
	if err := db.Where("id = ?", id).Delete(&Record{}).Error; err != nil {
		return err
	}
	return nil
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