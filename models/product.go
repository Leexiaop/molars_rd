package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Product struct {
	Model

	Name  string `json:"name"`
	Price int    `json:"price"`
	Url   string `json:"url"`
	CreatedBy   string `json:"created_by"`
	ModifieldBy string `json:"modifield_by"`
}

func GetProducts(pageNum int, pageSize int, maps interface{}) (products []Product) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&products)
	return
}

func GetProductsTotal(maps interface{}) (count int) {
	db.Model(&Product{}).Where(maps).Count(&count)
	return
}

func ExistProductName(name string) bool {
	var product Product
	db.Select("id").Where("name = ?", name).First(&product)
	return product.ID > 0
}

func ExistProductId (id int) bool {
	var product Product
	db.Select("id").Where("id = ?", id).First(&product)
	return product.ID > 0
}

func AddProducts(name string, price int, url string, createdBy string) bool {
	db.Create(&Product{
		Name:  name,
		Price: price,
		Url:   url,
		CreatedBy: createdBy,
	})
	return true
}

func EditProducts (id int, data interface{}) bool {
	db.Model(&Product{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteProducts (id int) bool {
	db.Where("id = ?", id).Delete(&Product{})
	return true
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (product *Product) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifieldOn", time.Now().Unix())
	return nil
}