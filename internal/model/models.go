package model

type Donation struct {
	ID         int    `json:"id"`
	FoodType   string `json:"food_type"`
	Quantity   int    `json:"quantity"`
	Expiration string `json:"expiration"`
	Location   string `json:"location"`
}
