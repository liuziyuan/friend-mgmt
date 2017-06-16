package dtos

// ConnectInput The connect friends input params
type ConnectInput struct {
	Friends [2]string `json:"friends"`
}

//RetrieveInput The retrieve friends input params
type RetrieveInput struct {
	Email string `json:"email"`
}
