package domain

// JWT describes a model
type JWT struct {
	Token     string `json:"token"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"`
}
