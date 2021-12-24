package model

type DataOwnedOrNotUserGroup struct {
	Users  DataOwnedOrNotUser  `json:"users"`
	Groups DataOwnedOrNotGroup `json:"groups"`
}
type DataOwnedOrNotUser struct {
	OwnedSlice    []User `json:"owned_slice"`
	NotOwnedSlice []User `json:"not_owned_slice"`
}
type DataOwnedOrNotGroup struct {
	OwnedSlice    []Group `json:"owned_slice"`
	NotOwnedSlice []Group `json:"not_owned_slice"`
}
type DataOwnedOrNotCluster struct {
	OwnedSlice    []Cluster `json:"owned_slice"`
	NotOwnedSlice []Cluster `json:"not_owned_slice"`
}
type DataOwnedOrNotRole struct {
	OwnedSlice    []Role `json:"owned_slice"`
	NotOwnedSlice []Role `json:"not_owned_slice"`
}
type DataOwnedOrNotPermission struct {
	OwnedSlice    []Permission `json:"owned_slice"`
	NotOwnedSlice []Permission `json:"not_owned_slice"`
}
