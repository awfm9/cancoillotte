// (c) 2019 Max Wolter - ALL RIGHTS RESERVED

package website

type Wizard struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	OwnerAddress string `json:"ownerAddress"`
	AffinityType string `json:"affinityType"`
	ImageURL     string `json:"imageUrl"`
	Power        string `json:"power"`
	Status       string `json:"status"`
	Owner        Owner  `json:"owner"`
}
