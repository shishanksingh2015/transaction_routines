package dao

type AccountDao struct {
	BaseDao
	Id             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}
