package payment

type PaymentStatus string

const (
	PaymentStatusPending  PaymentStatus = "PENDING"
	PaymentStatusApproved PaymentStatus = "APPROVED"
	PaymentStatusFailed   PaymentStatus = "FAILED"
)
