package models

type ACStatus struct {
	ACNumber int `json:"ac_number"` // AC number
	Status   int `json:"status"`    // 0 for off, 1 for on
}
type ACRemoteDTO struct {
	ACStatuses []ACStatus `json:"ac_statuses"` // List of AC statuses to control
}
