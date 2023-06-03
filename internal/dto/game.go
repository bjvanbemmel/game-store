package dto

type Game struct {
	id         int
	title      string
	price      float32
	developers []Developer
	categories []Category
}
