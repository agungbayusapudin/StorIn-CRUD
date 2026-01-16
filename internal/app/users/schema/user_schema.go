package schema

type Users struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	Phone           string `json:"phone"`
	Created_at      string `json:"created_at"`
	Created_at_unix int    `json:"created_at_unix"`
	Updated_at      string `json:"updated_at"`
	Updated_at_unix int    `json:"updated_at_unix"`
}

type UserRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
	Phone     *string `json:"phone"`
}
