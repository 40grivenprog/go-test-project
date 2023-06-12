package domain

// User belong to the domain layer.
type User struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	Email    string      `json:"email" bson:"email,omitempty"  binding:"required"`
	Password string      `json:"password" bson:"password,omitempty"  binding:"required"`
}
