package storage

// Resource is the struct used to store crawled resources
type Resource struct {
	ID                string `json:"id"`
	Name              string `json:"n"`
	ResourceGroupID   string `json:"rgid"`
	ResourceGroupName string `json:"rgn"`
}
