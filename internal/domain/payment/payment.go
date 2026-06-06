package payment

type Payment struct {
	Id        int
	BookingId int
	Amount    float64
	Status    PaymentStatus
}
