package payment

type PaymentModel struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	BookingID string  `gorm:"column:booking_id;not null"`
	Amount    float64 `gorm:"type:decimal(10,2)"`
	Status    string  `gorm:"column:status;size:20"`
}

