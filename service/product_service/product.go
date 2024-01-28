package product_service

import (
	"encoding/json"

	"github.com/Leexiaop/molars_rd/models"
	"github.com/Leexiaop/molars_rd/pkg/gredis"
	"github.com/Leexiaop/molars_rd/pkg/logging"
	"github.com/Leexiaop/molars_rd/service/cache_service"
)

type Product struct {
	ID         int
	Name       string
	Price      int
	Url        string
	CreatedBy  string
	ModifiedBy string

	PageSize int
	PageNum  int
}

func (p *Product) GetAll() ([]models.Product, error) {
	var (
		products, cachProducts []models.Product
	)

	cache := cache_service.Product{
		PageNum: p.PageNum,
		PageSize: p.PageSize,
	}

	key := cache.GetProductsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cachProducts)
			return cachProducts, nil
		}
	}
	products, err := models.GetProducts(p.PageNum, p.PageSize, p.getMaps())
	if err != nil {
		return nil, err
	}
	gredis.Set(key, products, 3600)
	return products, nil
}

func (p *Product) Add() error {
	return models.AddProducts(p.Name, p.Price, p.Url)
}
func (p *Product) ExistByName() (bool, error) {
	return models.ExistProductName(p.Name)
}
func(p *Product) Edit() error {
	data := make(map[string]interface{})
	if p.Name != "" {
		data["name"] = p.Name
	}
	if p.Price != 0 {
		data["price"] = p.Price
	}
	if p.Url != "" {
		data["url"] = p.Url
	}
	return models.EditProducts(p.ID, data)
}
func (p *Product) Count() (int, error) {
	return models.GetProductsTotal(p.getMaps())
}

func (p *Product) ExistById() (bool, error) {
	return models.ExistProductId(p.ID)
}

func (p *Product) Delete() error {
	return models.DeleteProduct(p.ID)
}

func(p*Product) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	return maps
}
