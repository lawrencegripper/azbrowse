package storage

type Resource struct {
	ID                string `json:"id"`
	Name              string `json:"n"`
	ResourceGroupID   string `json:"rgid"`
	ResourceGroupName string `json:"rgn"`
}
