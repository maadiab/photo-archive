package core

type User struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Username    string `db:"username"`
	Email       string `db:"email"`
	Mobile      string `db:"mobile"`
	Password    string `db:"hashedpassword"`
	Permissions string `db:"permission_type"`
}
