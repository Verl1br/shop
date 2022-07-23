package shop

type Item struct {
	Id      int    `json:"-" db:"id"`
	Name    string `json:"name" binding:"required"`
	Price   int    `json:"price" binding:"required"`
	BrandId int    `json:"brandid" binding:"required"`
}

type ItemUpdateInput struct {
	Name    *string `json:"name" binding:"required"`
	Price   *int    `json:"price" binding:"required"`
	BrandId *int    `json:"brandid" binding:"required"`
}

type ItemInfo struct {
	Id          int    `json:"-" db:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ItemId      int    `json:"itemid" binding:"required"`
}

type ItemInfoUpdateInput struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description" binding:"required"`
	ItemId      *int    `json:"itemid" binding:"required"`
}
