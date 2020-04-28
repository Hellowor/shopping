package model

type Product struct {
	ID           int64  `json:"id"`
	ProductName  string `json:"productName"`
	ProductNum   int64  `json:"productNum"`
	ProductImage string `json:"productImage"`
	ProductUrl   string `json:"productUrl"`
}

type IProduct interface {
	Conn() error
	Insert(product *Product)(int64,error)
	Delete(int64)bool
	Update(product *Product)error
	SelectByKey(int64)(*Product,error)
	SelectAll()
}