package dtos

//SubscriptionInput ...
type SubscriptionInput struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

//SubRetrieveInput ...
type SubRetrieveInput struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
