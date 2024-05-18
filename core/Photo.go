package core

type Photo struct {
	ID           int      `db:"id"`
	Name         string   `db:"name"`
	Photographer string   `db:"photographer"`
	Tags         []string `db:"tags"`
}
