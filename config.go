package main

type ShowInfo struct {
	ShowId     string
	SessionId  string
	SeatPlanId string
}

type AuthInfo struct {
	Token string
}

type AudienceInfo struct {
	BuyCount       int
	AudienceIndex  []int
	DeliveryMethod string
}
