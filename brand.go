package shop

type Brand struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}
