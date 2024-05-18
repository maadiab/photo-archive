package core

type Photographer struct {
	ID                   int    `db:"id"`
	PhotographerName     string `db:"name"`
	PhotographerUserName string `db:"username"`
	HashedPassword       string `db:"password"`
}
