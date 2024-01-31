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
	ModifieldBy string `json:"modifield_by"`
	CreatedBy string `json:"created_by"`
}

func GetProducts(pageNum int, pageSize int, maps interface{}) ([]Product, error) {
	var (
		products []Product
		err error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Offset(pageNum).Limit(pageSize).Where(maps).Find(&products).Error
	} else {
		err = db.Where(maps).Find(&products).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return products, nil
}

func GetProductsTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Product{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func ExistProductName(name string) (bool, error) {
	var product Product
	err := db.Select("id").Where("name = ?", name).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if product.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistProductId (id int) (bool, error) {
	var product Product
	err := db.Select("id").Where("id = ?", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if product.ID > 0 {
		return true, nil
	}
	return false, nil
}
func ExitRecords(productId int) (count int) {
	db.Model(&Record{}).Where("product_id = ?", productId).Count(&count)
	return
}

func AddProducts(name string, price int, url string, created_by string) error {
	product := Product{
		Name:  name,
		Price: price,
		Url:   url,
		CreatedBy: created_by,
	}
	if err := db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func EditProducts (id int, data interface{}) error {
	if err := db.Model(&Product{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProduct (id int) error {
	if err := db.Where("id = ?", id).Delete(&Product{}).Error; err != nil {
		return err
	}
	return nil
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (product *Product) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifieldOn", time.Now().Unix())
	return nil
}