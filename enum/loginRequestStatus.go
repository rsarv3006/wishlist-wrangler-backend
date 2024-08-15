package enum

const (
	LoginRequestStatusPending   = "PENDING"
	LoginRequestStatusCompleted = "COMPLETED"
	LoginRequestStatusExpired   = "EXPIRED"
)

var LoginRequestStatus = []string{
	LoginRequestStatusPending,
	LoginRequestStatusCompleted,
	LoginRequestStatusExpired,
}
