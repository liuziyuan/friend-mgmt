package dtos

//SubscriptionInput ...
type SubscriptionInput struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}
