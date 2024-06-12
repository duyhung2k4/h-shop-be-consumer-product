package model

type CheckcountPayload struct {
	ProductId         string
	WarehouseId       uint
	TypeInWarehouseId *uint
	GroupOrderId      uint
	Amount            uint
}
