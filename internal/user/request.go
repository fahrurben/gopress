package user

type (
	CreateUserRequest struct {
		Email    string `db:"email" json:"email"`
		Name     string `db:"name" json:"name"`
		Password string `db:"password"`
		Type     int    `db:"type" json:"type"`
	}
)
