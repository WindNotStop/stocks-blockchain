package model

const (
	A = iota
	HK
	US
)

type Stock struct {
	Uid    string `json:"stock_uid"`
	Name   string `json:"stock_name"`
	Date   string `json:"ipo_date"`
	Type   uint8  `json:"stock_type"`
	Amount uint64  `json:"stock_amount"`
	Price  float64 `json:"stock_price"`
}

func NewStock(uid, name, date string, Type uint8,amount uint64,price float64)*Stock{
	return &Stock{
		Uid:    uid,
		Name:   name,
		Date:   date,
		Type:   Type,
		Amount: amount,
		Price:  price,
	}
}
