package models

type ACStatus struct {
	ACNumber    int `json:"ac_number"`   // AC number
	Status      int `json:"status"`      // 0 for off, 1 for on
	Temperature int `json:"temperature"` // Current temperature setting
}

type ACRemoteRequest struct {
	ACStatuses []ACStatus `json:"ac_statuses"` // List of AC statuses to control
}

type ACRemoteResponse struct {
	ACStatuses []ACStatus `json:"ac_statuses"` // List of AC statuses after processing the request
}
