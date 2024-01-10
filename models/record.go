package models

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

func AddRecords () {}

func ExistRecordName () {}