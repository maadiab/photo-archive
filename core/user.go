package core

type User struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Username    string `db:"username" json:"username"`
	Email       string `db:"email" json:"email"`
	Mobile      string `db:"mobile" json:"mobile"`
	Password    string `db:"hashedpassword" json:"hashedpaassword"`
	Permissions string `db:"permission_type" json:"permission_type"`
}
