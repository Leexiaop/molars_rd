package models

type Product struct {
	Model

	Name  string `json:"name"`
	Price int    `json:"price"`
	Url   string `json:"url"`
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

func AddProducts(name string, price int, url string) bool {
	db.Create(&Product{
		Name:  name,
		Price: price,
		Url:   url,
	})
	return true
}