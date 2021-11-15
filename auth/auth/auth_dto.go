package auth

type AuthUserDTO struct {
	ID       string `db:"id,omitempty"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
