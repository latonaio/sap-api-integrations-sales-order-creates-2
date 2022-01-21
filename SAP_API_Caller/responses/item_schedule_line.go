package responses

type ItemScheduleLine struct {
	D struct {
	SalesOrder                     string `json:"SalesOrder"`
	SalesOrderItem                 string `json:"SalesOrderItem"`
	ScheduleLine                   string `json:"ScheduleLine"`
	RequestedDeliveryDate          string `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate          string `json:"ConfirmedDeliveryDate"`
	OrderQuantityUnit              string `json:"OrderQuantityUnit"`
	ScheduleLineOrderQuantity      string `json:"ScheduleLineOrderQuantity"`
	ConfdOrderQtyByMatlAvailCheck  string `json:"ConfdOrderQtyByMatlAvailCheck"`
	DeliveredQtyInOrderQtyUnit     string `json:"DeliveredQtyInOrderQtyUnit"`
	OpenConfdDelivQtyInOrdQtyUnit  string `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	CorrectedQtyInOrderQtyUnit     string `json:"CorrectedQtyInOrderQtyUnit"`
	DelivBlockReasonForSchedLine   string `json:"DelivBlockReasonForSchedLine"`
	} `json:"d"`
}
