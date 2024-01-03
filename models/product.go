package models

type Product struct {
	Model

	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	CreatedBy string `json:"created_by"`
	CreatedOn string `json:"created_on"`
	ModifieldBy string `json:"modifield_by"`
	ModifieldOn string `json:"modifield_on"`
	Url string `json:"url"`
}


func GetProducts (pageNum int, pageSize int, maps interface{}) (products []Product) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&products)
	return
}

func GetProductsTotal (maps interface{}) (count int) {
	db.Model(&Product{}).Where(maps).Count(&count)
	return
}