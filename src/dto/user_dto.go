package dto

type CreateUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	SocialNumber string `json:"social_number"`
}

type CreateUserResponse struct {
	CreatedUUID     string `json:"created_uuid"`
	CreatedUsername string `json:"created_username"`
	ResourceURL     string `json:"resource_url"`
}

type UserDetails struct {
	UserID       string `json:"user_id"`
	Username     string `json:"user_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	SocialNumber string `json:"social_number"`
}