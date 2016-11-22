package models

type Booking struct {
	IdBooking   int64 `db:"id_booking, primarykey, autoincrement"`
	ListService []*AgentService
	PriceTotal  float32 `db:"priceTotal"`
}
