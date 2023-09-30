package constants

const (
	// User Role
	UserRoleRegistrant int = 0
	UserRoleAdmin      int = 1

	// Registrant Status
	RegistrantStatusPending  int = 0
	RegistrantStatusOnReview int = 1
	RegistrantStatusAccepted int = 2
	RegistrantStatusRejected int = -1
)
