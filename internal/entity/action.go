package entity

type Action string

const (
	SEND                           Action = "SEND"
	SEND_WITHOUT_DELIVERED_STATE   Action = "SEND_WITHOUT_DELIVERED_STATE"
	DONT_SEND                      Action = "DONT_SEND"
	DONT_SEND_WITH_DELIVERED_STATE Action = "DONT_SEND_WITH_DELIVERED_STATE"
)
