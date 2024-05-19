package core

type Photographer struct {
	ID                   int    `db:"id" json:"id"`
	PhotographerName     string `db:"name" json:"name"`
	PhotographerUserName string `db:"username" json:"username"`
	HashedPassword       string `db:"password" json:"password"`
}
