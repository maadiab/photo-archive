package core

type Photo struct {
	ID           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	Photographer string `db:"photographer" json:"photographer"`
	Tags         string `db:"tags" json:"tags"`
}
